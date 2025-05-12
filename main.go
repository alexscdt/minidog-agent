package main

import (
	"fmt"
	"github.com/alexscdt/minidog-agent/internal/collector"
	"github.com/alexscdt/minidog-agent/internal/infrastructure/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("example.config.yaml")
	if err != nil {
		log.Fatalln("Error loading config: %v\n", err)
	}

	fmt.Println(cfg.Metrics.EnabledRAM)

	if cfg.Metrics.EnabledCPU {
		fmt.Println("CPU metrics: enabled")
		cpuMetrics := collector.CollectCpu()
		fmt.Printf("CPU Model: %s\n", cpuMetrics.Model)
	} else {
		fmt.Println("CPU metrics: disabled")
	}

	if cfg.Metrics.EnabledRAM {
		fmt.Println("RAM metrics: enabled")
		ramMetrics := collector.CollectRAM()
		fmt.Printf("RAM Total: %d\n", ramMetrics.Total)
	} else {
		fmt.Println("RAM metrics: disabled")
	}

	if cfg.Metrics.EnabledDisk {
		fmt.Println("Disk metrics: enabled")
		diskMetrics := collector.CollectDisk()
		for _, disk := range diskMetrics {
			fmt.Printf("Disk Path: %s, Total: %d\n", disk.Path, disk.Total)
		}
	} else {
		fmt.Println("Disk metrics: disabled")
	}

	if cfg.Metrics.EnabledNetwork {
		fmt.Println("Network metrics: enabled")
		networkMetrics := collector.CollectNetwork()
		for _, network := range networkMetrics {
			fmt.Printf("Network Interface: %s, Bytes Sent: %d\n", network.InterfaceName, network.BytesSent)
		}
	} else {
		fmt.Println("Network metrics: disabled")
	}

}
