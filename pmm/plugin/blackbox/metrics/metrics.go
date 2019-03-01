package metrics

import (
	"context"
	"fmt"

	"github.com/percona/pmm-client/pmm"
	"github.com/percona/pmm-client/pmm/plugin"
	"github.com/percona/pmm-client/pmm/plugin/blackbox"
)

var _ plugin.Metrics = (*Metrics)(nil)

// New returns *Metrics.
func New(flags blackbox.Flags) *Metrics {
	return &Metrics{}
}

// Metrics implements plugin.Metrics.
type Metrics struct{}

// Init initializes plugin.
func (m *Metrics) Init(ctx context.Context, pmmUserPassword string) (*plugin.Info, error) {
	info := &plugin.Info{}
	return info, nil
}

// Name of the exporter.
func (Metrics) Name() string {
	return "blackbox"
}

// DefaultPort returns default port.
func (Metrics) DefaultPort() int {
	return 42012
}

// Args is a list of additional arguments passed to exporter executable.
func (Metrics) Args() []string {
	return []string{
		fmt.Sprintf("--config.file=%s/blackbox.yml", pmm.PMMBaseDir),
	}
}

// Environment is a list of additional environment variables passed to exporter executable.
func (m Metrics) Environment() []string {
	return []string{}
}

// Executable is a name of exporter executable under PMMBaseDir.
func (Metrics) Executable() string {
	return "blackbox_exporter"
}

// KV is a list of additional Key-Value data stored in consul.
func (m Metrics) KV() map[string][]byte {
	return map[string][]byte{}
}

// Cluster defines cluster name for the target.
func (Metrics) Cluster() string {
	return ""
}

// Multiple returns true if exporter can be added multiple times.
func (Metrics) Multiple() bool {
	return true
}
