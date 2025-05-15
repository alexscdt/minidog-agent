package application

import (
	"fmt"
	"github.com/alexscdt/minidog-agent/internal/collector"
	"github.com/alexscdt/minidog-agent/internal/domain"
	"github.com/alexscdt/minidog-agent/internal/infrastructure/config"
	"github.com/alexscdt/minidog-agent/internal/infrastructure/logwatcher"
	"log"
	"time"
)

func StartAgent(cfg *config.Config) {

	fmt.Println("Starting agent...")

	fmt.Printf("Agent interval: %d seconds\n", cfg.Agent.Interval)
	fmt.Printf("Agent id: %s\n", cfg.Agent.ID)

	fmt.Println("Activate metrics...")
	fmt.Printf("CPU metrics: %t\n", cfg.Metrics.EnabledCPU)
	fmt.Printf("RAM metrics: %t\n", cfg.Metrics.EnabledRAM)
	fmt.Printf("Disk metrics: %t\n", cfg.Metrics.EnabledDisk)
	fmt.Printf("Network metrics: %t\n", cfg.Metrics.EnabledNetwork)

	fmt.Println("Path to log file: ", cfg.Logs.Paths)

	ticker := time.NewTicker(time.Duration(cfg.Agent.Interval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if cfg.Metrics.EnabledCPU {
				cpuMetrics := collector.CollectCpu()
				fmt.Printf("CPU Model: %s\n", cpuMetrics.Model)
			}

			if cfg.Metrics.EnabledRAM {
				ramMetrics := collector.CollectRAM()
				fmt.Printf("RAM Total: %d\n", ramMetrics.Total)
			}

			if cfg.Metrics.EnabledDisk {
				diskMetrics := collector.CollectDisk()
				for _, disk := range diskMetrics {
					fmt.Printf("Disk Path: %s, Total: %d\n", disk.Path, disk.Total)
				}
			}

			if cfg.Metrics.EnabledNetwork {
				networkMetrics := collector.CollectNetwork()
				for _, network := range networkMetrics {
					fmt.Printf("Network Interface: %s, Bytes Sent: %d\n", network.InterfaceName, network.BytesSent)
				}
			}

			logChannel := make(chan domain.LogEvent, 100)

			if cfg.Metrics.EnableLogs {
				fmt.Println("Starting log watchers...")
				for _, path := range cfg.Logs.Paths {
					go logwatcher.WatchFile(path, logChannel)
					log.Printf("Watcher lancé pour %s", path)
				}

				go func() {
					for logEvent := range logChannel {
						fmt.Printf("%s | %s : [%s] %s\n", logEvent.Timestamp, logEvent.Level, logEvent.FilePath, logEvent.Content)
					}
				}()

			}
		}
	}
}
