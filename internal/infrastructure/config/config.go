package config

type AgentConfig struct {
	ID         string `yaml:"id"`
	Token      string `yaml:"token"`
	Interval   int    `yaml:"interval"`
	BackendUrl string `yaml:"backend_url"`
}

type MetricsConfig struct {
	EnabledCPU     bool `yaml:"enable_cpu"`
	EnabledRAM     bool `yaml:"enable_memory"`
	EnabledDisk    bool `yaml:"enable_disk"`
	EnabledNetwork bool `yaml:"enable_network"`
	EnableLogs     bool `yaml:"enable_logs"`
}

type LogsConfig struct {
	Paths []string `yaml:"paths"`
}

type Config struct {
	Agent   AgentConfig   `yaml:"agent"`
	Metrics MetricsConfig `yaml:"metrics"`
	Logs    LogsConfig    `yaml:"logs"`
}
