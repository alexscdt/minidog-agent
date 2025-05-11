package config

type AgentConfig struct {
	ID         string `yaml:"id"`
	Token      string `yaml:"token"`
	Interval   int    `yaml:"interval"`
	BackendUrl string `yaml:"backend_url"`
}

type MetricsConfig struct {
	EnabledCPU     bool `yaml:"enabledCPU"`
	EnabledRAM     bool `yaml:"enabledRAM"`
	EnabledDisk    bool `yaml:"enabledDisk"`
	EnabledNetwork bool `yaml:"enabledNetwork"`
}

type LogsConfig struct {
	Paths []string `yaml:"paths"`
}

type Config struct {
	Agent   AgentConfig   `yaml:"agent"`
	Metrics MetricsConfig `yaml:"metrics"`
	Logs    LogsConfig    `yaml:"logs"`
}
