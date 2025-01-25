package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var config Config
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("problem with getting the config file path %v", err)
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return Config{}, err

	}
	return config, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("problem with getting the home dir %v", err)
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)

	if err != nil {
		return err
	}
	return nil
}
