package main

import (
	"bixpark_server/config"
	"bixpark_server/internal/data"
	"flag"
	"log"
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
	// Load App configs
	log.Println("Load config success ", config.App.Name)

	//dB
	db, err := DBConnect(config)
	if err != nil {
		log.Println("Error in DB Connect")
	}
	defer db.Close()
	data.SetupDB(db)

}
