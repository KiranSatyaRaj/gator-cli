package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_URL          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", homeDir, configFileName), nil
}

func Read() Config {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	defer configFile.Close()

	return config
}

func (config *Config) SetUser() {
	config.CurrentUserName = os.Getenv("USER")
	if err := write(*config); err != nil {
		panic(err)
	}
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(&cfg)
	if err != nil {
		panic(err)
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(configFilePath, jsonData, 0o755)
	return err
}
