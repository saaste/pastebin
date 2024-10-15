package config

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const appVersion string = "0.1.0"

type AppConfig struct {
	BaseURL       string `yaml:"baseUrl"`
	Title         string `yaml:"title"`
	Theme         string `yaml:"theme"`
	Password      string `yaml:"password"`
	JwtSecret     string `yaml:"jwtSecret"`
	DefaultSyntax string `yaml:"defaultSyntax"`
	AppVersion    string
}

func GetConfig() (*AppConfig, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	appConfig := &AppConfig{
		AppVersion: appVersion,
	}

	err = yaml.Unmarshal(yamlFile, appConfig)
	if err != nil {
		return nil, err
	}

	appConfig.BaseURL = strings.TrimSuffix(appConfig.BaseURL, "/")

	return appConfig, nil
}
