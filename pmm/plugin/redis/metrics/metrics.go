package metrics

import (
	"context"
	"fmt"

	redisgo "github.com/gomodule/redigo/redis"
	"github.com/percona/pmm-client/pmm/plugin"
	"github.com/percona/pmm-client/pmm/plugin/redis"
)

var _ plugin.Metrics = (*Metrics)(nil)

// New returns *Metrics.
func New(flags redis.Flags) *Metrics {
	return &Metrics{
		url: flags.Url,
	}
}

// Metrics implements plugin.Metrics.
type Metrics struct {
	url string
}

// Init initializes plugin.
func (m *Metrics) Init(ctx context.Context, pmmUserPassword string) (*plugin.Info, error) {
	if err := testConnection(m.url); err != nil {
		return nil, fmt.Errorf("cannot connect to Redis using Address %s: %s", m.url, err)
	}

	info := &plugin.Info{
		DSN: m.url,
	}
	return info, nil
}

// Name of the exporter.
func (Metrics) Name() string {
	return "redis"
}

// DefaultPort returns default port.
func (Metrics) DefaultPort() int {
	return 30000
}

// Args is a list of additional arguments passed to exporter executable.
func (Metrics) Args() []string {
	return nil
}

// Environment is a list of additional environment variables passed to exporter executable.
func (m Metrics) Environment() []string {
	return []string{
		fmt.Sprintf("REDIS_URL=%s", m.url),
	}
}

// Executable is a name of exporter executable under PMMBaseDir.
func (Metrics) Executable() string {
	return "redis_exporter"
}

// KV is a list of additional Key-Value data stored in consul.
func (m Metrics) KV() map[string][]byte {
	return map[string][]byte{
		"url": []byte(m.url),
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

func testConnection(url string) error {
	c, err := redisgo.DialURL(url)
	if err != nil {
		return err
	}
	defer c.Close()
	_, err = c.Do("PING")
	return err
}
