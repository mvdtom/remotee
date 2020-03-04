package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Builder struct {
		TemplateFile string `yaml:"template_file"`
		Offset       int    `yaml:"offset"`
		StorageDir   string `yaml:"storage_dir"`
	}
}

func GetConfig() (*Config, error) {
	f, err := os.Open("config/config.yml")
	if err != nil {
		fmt.Println("Error opening config file", err.Error())
		return nil, err
	}
	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Error reading config", err.Error())
		os.Exit(1)
	}
	return &cfg, nil
}
