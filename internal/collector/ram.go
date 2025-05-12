package collector

import (
	"log"

	"github.com/alexscdt/minidog-agent/internal/domain"
	"github.com/shirou/gopsutil/v3/mem"
)

func CollectRAM() domain.RamMetrics {
	vm, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("RAM error: %v", err)
		return domain.RamMetrics{}
	}

	sm, err := mem.SwapMemory()
	if err != nil {
		log.Printf("Swap error: %v", err)
	}

	return domain.RamMetrics{
		Total:        vm.Total,
		Used:         vm.Used,
		Free:         vm.Free,
		Available:    vm.Available,
		UsedPercent:  vm.UsedPercent,
		SwapTotal:    sm.Total,
		SwapUsed:     sm.Used,
		SwapUsedPerc: sm.UsedPercent,
	}
	
}
