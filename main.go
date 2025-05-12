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
		collector.CollectCpu()
	} else {
		fmt.Println("CPU metrics: disabled")
	}

	if cfg.Metrics.EnabledRAM {
		fmt.Println("RAM metrics: enabled")
	} else {
		fmt.Println("RAM metrics: disabled")
	}

	if cfg.Metrics.EnabledDisk {
		fmt.Println("Disk metrics: enabled")
	} else {
		fmt.Println("Disk metrics: disabled")
	}

	if cfg.Metrics.EnabledNetwork {
		fmt.Println("Network metrics: enabled")
	} else {
		fmt.Println("Network metrics: disabled")
	}

}
