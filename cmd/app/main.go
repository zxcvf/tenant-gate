package main

import (
	"log"
	"tenant-gate/config"
	"tenant-gate/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	log.Printf("Config: %+v", cfg)

	app.Run(cfg)
}
