package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DBurl       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	var configPath string

	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath = homePath + configFileName

	return configPath, nil
}

func Read() (Config, error) {
	var configContents Config

	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	cont, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(cont, &configContents)

	return configContents, nil
}

func write(cfg Config) error {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, bytes, 0644)

	return nil
}

func (c Config) SetUser(user_name string) error {
	c.CurrentUser = user_name

	err := write(c)
	if err != nil {
		return err
	}

	return nil
}

//
