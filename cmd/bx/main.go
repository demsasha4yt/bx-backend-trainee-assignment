package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/bx"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/emailsender"
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

func parseSMTPEnvironment(config *emailsender.Config) {
	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost != "" {
		config.ServerHost = smtpHost
	}
	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort != "" {
		config.ServerPort = smtpPort
	}
	smtpUser := os.Getenv("SMTP_USER")
	if smtpUser != "" {
		config.Username = smtpUser
	}
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	if smtpPassword != "" {
		config.Password = smtpPassword
	}
	smtpSender := os.Getenv("SMTP_SENDER_ADDR")
	if smtpSender != "" {
		config.SenderAddr = smtpSender
	}
}

func main() {
	flag.Parse()

	config := bx.NewConfig()
	configData, err := getConfigData(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatal(err)
		return
	}

	parseSMTPEnvironment(&config.SMTPConfig)

	if err := bx.Start(config); err != nil {
		log.Fatal(err)
		return
	}
}
