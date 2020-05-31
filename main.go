package main

import (
	"bixpark_server/config"
	"flag"
	"os"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", os.Getenv("BIX_PARK_APP_CONFIG"), "App config locations")
	flag.Parse()

	config, err := config.NewConfig(configPath)
	if err != nil {
		println("Error on reading config")
	}
	println(config.App.Name)
}
