package domain

import "time"

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

type RamMetrics struct {
	Total        uint64  `json:"total"`
	Used         uint64  `json:"used"`
	Free         uint64  `json:"free"`
	UsedPercent  float64 `json:"used_percent"`
	Available    uint64  `json:"available"`
	SwapTotal    uint64  `json:"swap_total"`
	SwapUsed     uint64  `json:"swap_used"`
	SwapUsedPerc float64 `json:"swap_used_percent"`
}

type DiskMetrics struct {
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
	ReadCount   uint64  `json:"read_count"`
	WriteCount  uint64  `json:"write_count"`
	ReadBytes   uint64  `json:"read_bytes"`
	WriteBytes  uint64  `json:"write_bytes"`
}

type NetworkMetrics struct {
	InterfaceName string `json:"interface_name"`
	BytesSent     uint64 `json:"bytes_sent"`
	BytesRecv     uint64 `json:"bytes_recv"`
	PacketsSent   uint64 `json:"packets_sent"`
	PacketsRecv   uint64 `json:"packets_recv"`
	ErrIn         uint64 `json:"err_in"`
	ErrOut        uint64 `json:"err_out"`
}

type MetricsPayload struct {
	AgentID   string           `json:"agent_id"`
	CPU       CPUMetrics       `json:"cpu"`
	RAM       RamMetrics       `json:"ram"`
	Disk      []DiskMetrics    `json:"disk"`
	Network   []NetworkMetrics `json:"network"`
	Timestamp time.Time        `json:"timestamp"`
}
