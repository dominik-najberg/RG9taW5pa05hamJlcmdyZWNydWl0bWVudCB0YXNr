package main

import (
	"fmt"
	"github.com/dominik-najberg/gogoapps/config"
	"github.com/dominik-najberg/gogoapps/internal"
	"github.com/dominik-najberg/gogoapps/server"
	"log"
	"os"
)

func main() {
	// just for fun
	if len(os.Args) > 1 && os.Args[1] == "github" {
		fmt.Println("GitHub name for GoGoApps:", internal.CalculateBase64())
		os.Exit(0)
	}

	log.Println("reading config")
	config.ReadConfig("config/config.json")

	log.Println("initializing cache")
	server.InitCache()

	log.Println("starting apiServer...")
	server.Start(config.AppConfig.ServerPort)
	log.Println("apiServer stopped")
}


