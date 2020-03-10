package config

import (
	"encoding/json"
	"log"
	"os"
)

var AppConfig Configuration

type Configuration struct {
	ServerPort int `json:"local_port"`
}

func ReadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error while opening config file: %v", err)
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("error while decoding configuration: %v", err)
	}
}
