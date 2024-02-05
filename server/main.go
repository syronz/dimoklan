package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"dimoklan/api"
	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/service"
	"dimoklan/storage"
)

var configFilePath = flag.String("cfg", "", "config file path")

func main() {
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
			fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
			fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
			fmt.Printf("\tNumGC = %v\n", m.NumGC)
			time.Sleep(10 * time.Second)
		}
	}()

	flag.Parse()

	if *configFilePath == "" {
		log.Fatal("cfg is required")
	}

	core, err := config.GetCore(*configFilePath)
	if err != nil {
		log.Fatalf("error in loading core; %v", err)
	}

	core.Info("starting server: " + time.Now().String())

	storage := storage.NewMemroryStorage()
	basStorage := basstorage.New(core)
	userService := service.NewUserService(core, basStorage)

	server := api.NewServer(core, storage, userService)
	server.Start()
}
