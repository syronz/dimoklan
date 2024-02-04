package main

import (
	"flag"
	"log"
	"time"

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

	core, err := config.GetCore(*configFilePath)
	if err != nil {
		log.Fatalf("error in loading core; %v", err)
	}

	core.Info(time.Now().String())

	storage := storage.NewMemroryStorage()
	basStorage := basstorage.New(core)

	server := api.NewServer(core, storage, basStorage)
	server.Start()
}
