package redis

// Flags are Orchestrator specific flags.
type Flags struct {
	Addr                    string
	File                    string
	Password                string
	PasswordFile            string
	Alias                   string
	ExporterNamespace       string
	ExporterCheckKeys       string
	ExporterCheckSingleKeys string
	ExporterScript          string
	ExporterSeparator       string
	ExporterDebug           string
	ExporterLogFormat       string
}
