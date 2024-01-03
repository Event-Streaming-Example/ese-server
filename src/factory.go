package main

import (
	"errors"
	"ese/server/application"
	"ese/server/data"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Redis  RedisConfig  `yaml:"redis"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Url  string `yaml:"url"`
}

type RedisConfig struct {
	Port            int    `yaml:"port"`
	Url             string `yaml:"url"`
	ExpiryInMinutes int    `yaml:"expiryInMinutes"`
}

func ProvideRedisClient(config RedisConfig) data.RedisClient {
	return data.ProvideRedisClient(config.Url, config.Port, config.ExpiryInMinutes)
}

func ProvideServer(config ServerConfig, redisClient *data.RedisClient) application.Server {
	address := fmt.Sprintf("%s:%d", config.Url, config.Port)
	return application.ProvideServer(redisClient, address)
}

func ProvideConfig() Config {
	file, err := os.Open("./properties/config.yaml")
	if err != nil {
		log.Fatal("Error reading YAML file: ", err)
		errors.New("Invalid Configs passed")
	}
	defer file.Close()

	// Read the YAML content
	var yamlContent []byte
	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading YAML file: ", err)
			errors.New("Invalid Configs passed")
		}
		yamlContent = append(yamlContent, buffer[:n]...)
	}

	// Unmarshal the YAML content into a struct
	var config Config
	err = yaml.Unmarshal(yamlContent, &config)
	if err != nil {
		log.Fatal("Error reading YAML file: ", err)
		errors.New("Invalid Configs passed")
	}
	return config
}
