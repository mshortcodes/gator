package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	err := write(cfg)
	if err != nil {
		return fmt.Errorf("couldn't set user")
	}
	return nil
}

func write(cfg *Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("couldn't marshal struct")
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("couldn't get file path")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("couldn't create file")
	}

	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("couldn't write to file")
	}

	return nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("couldn't read file")
	}

	var cfg Config
	err = json.Unmarshal(jsonData, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("couldn't unmarshal json")
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("couldn't get HOME env var")
	}
	filePath := home + "/" + configFileName
	return filePath, nil
}
