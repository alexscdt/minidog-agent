package collector

import (
	"fmt"
	"github.com/alexscdt/minidog-agent/internal/domain"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
)

func CollectCpu() {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		log.Fatalf("Failed to collect CPU metrics: %v", err)
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Fatalf("Failed to collect CPU info: %v", err)
	}

	logicalCount, err := cpu.Counts(true)
	if err != nil {
		log.Fatalf("Failed to get logical CPU count: %v", err)
	}

	physicalCount, err := cpu.Counts(false)
	if err != nil {
		log.Fatalf("Failed to get physical CPU count: %v", err)
	}

	cpuTimes, err := cpu.Times(false)
	if err != nil {
		log.Fatalf("Failed to get CPU times: %v", err)
	}

	cpuStat := domain.CPUMetrics{
		Model:        cpuInfo[0].ModelName,
		Cores:        physicalCount,
		Threads:      logicalCount,
		FrequencyGHz: float64(cpuInfo[0].Mhz) / 1000.0,
		UsagePercent: percentages[0],
		UserTime:     cpuTimes[0].User,
		SystemTime:   cpuTimes[0].System,
		IdleTime:     cpuTimes[0].Idle,
	}

	fmt.Printf("CPU info: %v\n", cpuStat)
}
