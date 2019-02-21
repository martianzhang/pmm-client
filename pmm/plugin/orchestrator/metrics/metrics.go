package metrics

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/percona/pmm-client/pmm/plugin"
	"github.com/percona/pmm-client/pmm/plugin/orchestrator"
)

var _ plugin.Metrics = (*Metrics)(nil)

// New returns *Metrics.
func New(flags orchestrator.Flags) *Metrics {
	return &Metrics{
		api:      flags.Api,
		user:     flags.User,
		password: flags.Password,
	}
}

// Metrics implements plugin.Metrics.
type Metrics struct {
	api      string
	user     string
	password string
}

// Init initializes plugin.
func (m *Metrics) Init(ctx context.Context, pmmUserPassword string) (*plugin.Info, error) {
	if err := testConnection(m.user, m.password, m.api); err != nil {
		return nil, fmt.Errorf("cannot connect to Orchestrator using API %s (%s:%s): %s", m.api, m.user, m.password, err)
	}

	info := &plugin.Info{
		DSN: m.api,
	}
	return info, nil
}

// Name of the exporter.
func (Metrics) Name() string {
	return "orchestrator"
}

// DefaultPort returns default port.
func (Metrics) DefaultPort() int {
	return 42010
}

// Args is a list of additional arguments passed to exporter executable.
func (Metrics) Args() []string {
	return nil
}

// Environment is a list of additional environment variables passed to exporter executable.
func (m Metrics) Environment() []string {
	return []string{
		fmt.Sprintf("ORCHESTRATOR_AUTH_SERVER=%s", m.api),
		fmt.Sprintf("ORCHESTRATOR_AUTH_USER=%s", m.user),
		fmt.Sprintf("ORCHESTRATOR_AUTH_PASSWORD=%s", m.password),
	}
}

// Executable is a name of exporter executable under PMMBaseDir.
func (Metrics) Executable() string {
	return "orchestrator_exporter"
}

// KV is a list of additional Key-Value data stored in consul.
func (m Metrics) KV() map[string][]byte {
	return map[string][]byte{
		"api": []byte(m.api),
	}
}

// Cluster defines cluster name for the target.
func (Metrics) Cluster() string {
	return ""
}

// Multiple returns true if exporter can be added multiple times.
func (Metrics) Multiple() bool {
	return true
}

// basicAuth user & password crypto
func basicAuth(user, password string) string {
	auth := user + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func testConnection(user, password, api string) error {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	var url string
	if strings.HasSuffix(api, "/") {
		url = api + "health"
	} else {
		url = api + "/health"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if user != "" {
		req.Header.Add("Authorization", "Basic "+basicAuth(user, password))
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		resp.Body.Close()
		return fmt.Errorf("HTTP status %d", resp.StatusCode)
	}
	return nil
}
