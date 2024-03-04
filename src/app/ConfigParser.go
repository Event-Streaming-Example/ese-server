package app

import (
	"io"
	"log"
	"os"

	"ese.server/redis"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Redis redis.Config `yaml:"redis"`
}

const PROPERTIES_FILE_PATH = "./app/properties/config.yaml"
const ERROR_PARSING_CONFIG = "Error reading .yml file. Invaid configs passed : "

func ParseConfig() Config {
	file, err := os.Open(PROPERTIES_FILE_PATH)
	if err != nil {
		log.Fatal(ERROR_PARSING_CONFIG, err)
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
			log.Fatal(ERROR_PARSING_CONFIG, err)
		}
		yamlContent = append(yamlContent, buffer[:n]...)
	}

	// Unmarshal the YAML content into a struct
	var config Config
	err = yaml.Unmarshal(yamlContent, &config)
	if err != nil {
		log.Fatal(ERROR_PARSING_CONFIG, err)
	}
	return config
}
