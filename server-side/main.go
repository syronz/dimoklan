package main

import (
	"flag"
	"log"

	"dimoklan/api"
	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/storage"
)

var configFilePath = flag.String("config-path", "", "config file path")

func main() {
	flag.Parse()

	if *configFilePath == "" {
		log.Fatal("config-path is required")
	}

	cfg, err := config.GetConf(*configFilePath)
	if err != nil {
		log.Fatalf("error in loading config; %v", err)
	}

	storage := storage.NewMemroryStorage()
	basStorage := basstorage.New(cfg)

	server := api.NewServer(cfg, storage, basStorage)
	server.Start()
}
