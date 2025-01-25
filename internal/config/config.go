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
	defer file.Close()

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
	// Create temp file for atomic write
	tmpFile, err := os.CreateTemp(filepath.Dir(configFilePath), "*.tmp")

	if err != nil {
		return err
	}

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	encoder := json.NewEncoder(tmpFile)
	err = encoder.Encode(cfg)

	if err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	// Atomic rename
	if err := os.Rename(tmpFile.Name(), configFilePath); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}
