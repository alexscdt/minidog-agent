package main

import (
	"fmt"
	"github.com/alexscdt/minidog-agent/internal/infrastructure/config"
	"log"
)

func main() {
	fmt.Println("Hello, World!")

	cfg, err := config.LoadConfig("example.config.yaml")
	if err != nil {
		log.Fatalln("Error loading config: %v\n", err)
	}

	fmt.Printf("Agent ID: %s\n", cfg.Agent.ID)

}
