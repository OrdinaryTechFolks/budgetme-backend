package main

import (
	"flag"

	cfg "github.com/OrdinaryTechFolks/budgetme-backend/internal/config"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./files/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := cfg.GetConfig(configPath)
	if err != nil {
		panic(err)
	}

	startApp(cfg)
}
