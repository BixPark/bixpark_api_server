package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		MediaPath   string `yaml:"mediaPath"`
	}

	Route struct {
		Prefix  string `yaml:"prefix"`
		Version string `yaml:"version"`
	}

	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		BbName   string `yaml:"db"`
	}

	Security struct {
		Salt string `yaml:"salt"`
	}
}

func newConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func NewConfig(configPath string) (*Config, error) {
	// Validate the path first
	if err := validateConfigPath(configPath); err != nil {
		return nil, err
	}
	return newConfig(configPath)
}
