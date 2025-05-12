package collector

import (
	"log"

	"github.com/alexscdt/minidog-agent/internal/domain"
	"github.com/shirou/gopsutil/v3/disk"
)

func CollectDisk() []domain.DiskMetrics {
	usage, err := disk.Usage("/")
	if err != nil {
		log.Printf("Disk usage error: %v", err)
		return nil
	}

	ioStats, err := disk.IOCounters()
	if err != nil {
		log.Printf("Disk IO error: %v", err)
		return nil
	}

	var ioStat disk.IOCountersStat
	for _, stat := range ioStats {
		ioStat = stat
		break
	}

	return []domain.DiskMetrics{
		{
			Path:        usage.Path,
			Total:       usage.Total,
			Used:        usage.Used,
			Free:        usage.Free,
			UsedPercent: usage.UsedPercent,
			ReadCount:   ioStat.ReadCount,
			WriteCount:  ioStat.WriteCount,
			ReadBytes:   ioStat.ReadBytes,
			WriteBytes:  ioStat.WriteBytes,
		},
	}
}
