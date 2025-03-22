package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	NotificationURL string `yaml:"notification_url"`
}

func loadConfig() (*Config, error) {
	var config Config
	yamlFile, err := os.ReadFile("schedule.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func sendSlackMessage(uri string, message string) {
	slackMessage := fmt.Sprintf(`{"text": "%s"}`, message)

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(slackMessage)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	sendSlackMessage(config.NotificationURL, "Hello, World!")
}
