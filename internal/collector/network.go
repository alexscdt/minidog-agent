package collector

import (
	"log"

	"github.com/alexscdt/minidog-agent/internal/domain"
	"github.com/shirou/gopsutil/v3/net"
)

func CollectNetwork() []domain.NetworkMetrics {
	netStats, err := net.IOCounters(true)
	if err != nil {
		log.Printf("Network error: %v", err)
		return nil
	}

	var result []domain.NetworkMetrics
	for _, iface := range netStats {
		result = append(result, domain.NetworkMetrics{
			InterfaceName: iface.Name,
			BytesSent:     iface.BytesSent,
			BytesRecv:     iface.BytesRecv,
			PacketsSent:   iface.PacketsSent,
			PacketsRecv:   iface.PacketsRecv,
			ErrIn:         iface.Errin,
			ErrOut:        iface.Errout,
		})
	}

	return result
}
