package main

import (
	"github.com/alexscdt/minidog-agent/internal/application"
	"github.com/alexscdt/minidog-agent/internal/infrastructure/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("example.config.yaml")
	if err != nil {
		log.Fatalln("Error loading config: %v\n", err)
	}

	application.StartAgent(cfg)

}
