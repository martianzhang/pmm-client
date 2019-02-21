package metrics

import (
	"context"
	"fmt"

	"github.com/percona/pmm-client/pmm/plugin"
	"github.com/percona/pmm-client/pmm/plugin/redis"
)

var _ plugin.Metrics = (*Metrics)(nil)

// New returns *Metrics.
func New(flags redis.Flags) *Metrics {
	return &Metrics{
		addr:                    flags.Addr,
		file:                    flags.File,
		password:                flags.Password,
		passwordFile:            flags.PasswordFile,
		alias:                   flags.Alias,
		exporterNamespace:       flags.ExporterNamespace,
		exporterCheckKeys:       flags.ExporterCheckKeys,
		exporterCheckSingleKeys: flags.ExporterCheckSingleKeys,
		exporterScript:          flags.ExporterScript,
		exporterSeparator:       flags.ExporterSeparator,
		exporterDebug:           flags.ExporterDebug,
		exporterLogFormat:       flags.ExporterLogFormat,
	}
}

// Metrics implements plugin.Metrics.
type Metrics struct {
	addr                    string
	file                    string
	password                string
	passwordFile            string
	alias                   string
	exporterNamespace       string
	exporterCheckKeys       string
	exporterCheckSingleKeys string
	exporterScript          string
	exporterSeparator       string
	exporterDebug           string
	exporterLogFormat       string
}

// Init initializes plugin.
func (m *Metrics) Init(ctx context.Context, pmmUserPassword string) (*plugin.Info, error) {
	info := &plugin.Info{
		DSN: m.addr,
	}
	return info, nil
}

// Name of the exporter.
func (Metrics) Name() string {
	return "redis"
}

// DefaultPort returns default port.
func (Metrics) DefaultPort() int {
	return 42011
}

// Args is a list of additional arguments passed to exporter executable.
func (Metrics) Args() []string {
	return nil
}

// Environment is a list of additional environment variables passed to exporter executable.
func (m Metrics) Environment() []string {
	return []string{
		fmt.Sprintf("REDIS_ADDR=%s", m.addr),
		fmt.Sprintf("REDIS_FILE=%s", m.file),
		fmt.Sprintf("REDIS_PASSWORD=%s", m.password),
		fmt.Sprintf("REDIS_PASSWORD_FILE=%s", m.passwordFile),
		fmt.Sprintf("REDIS_ALIAS=%s", m.alias),
		fmt.Sprintf("REDIS_EXPORTER_NAMESPACE=%s", m.exporterNamespace),
		fmt.Sprintf("REDIS_EXPORTER_CHECK_KEYS=%s", m.exporterCheckKeys),
		fmt.Sprintf("REDIS_EXPORTER_CHECK_SINGLE_KEYS=%s", m.exporterCheckSingleKeys),
		fmt.Sprintf("REDIS_EXPORTER_SCRIPT=%s", m.exporterScript),
		fmt.Sprintf("REDIS_EXPORTER_SEPARATOR=%s", m.exporterSeparator),
		fmt.Sprintf("REDIS_EXPORTER_DEBUG=%s", m.exporterDebug),
		fmt.Sprintf("REDIS_EXPORTER_LOG_FORMAT=%s", m.exporterLogFormat),
	}
}

// Executable is a name of exporter executable under PMMBaseDir.
func (Metrics) Executable() string {
	return "redis_exporter"
}

// KV is a list of additional Key-Value data stored in consul.
func (m Metrics) KV() map[string][]byte {
	return map[string][]byte{
		"addr": []byte(m.addr),
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
