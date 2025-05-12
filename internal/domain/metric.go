package domain

type CPUMetrics struct {
	Model        string  `json:"model"`
	Cores        int     `json:"cores"`
	Threads      int     `json:"threads"`
	FrequencyGHz float64 `json:"frequency_ghz"`
	UsagePercent float64 `json:"usage_percent"`
	UserTime     float64 `json:"user_time"`
	SystemTime   float64 `json:"system_time"`
	IdleTime     float64 `json:"idle_time"`
}
