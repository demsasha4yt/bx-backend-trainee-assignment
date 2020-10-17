package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitomock"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/bx.json", "path to config file")
}

func getConfigData(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, err
}

func main() {
	flag.Parse()

	config := avitomock.NewConfig()
	configData, err := getConfigData(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatal(err)
		return
	}
	avitomock.Start(config)
}
