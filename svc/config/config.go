package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Api struct {
		Key string `yaml:"key"`
	} `yaml:"api"`
}

var appConfig Config

func LoadConfig() *Config {
	filePath := "config/config.yaml"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to read config file", err)
	}
	marshalErr := yaml.Unmarshal(data, &appConfig)
	if marshalErr != nil {
		log.Fatal("Failed to parse", marshalErr)
	}
	return &appConfig
}

func GetConfig() *Config {
	return &appConfig
}
