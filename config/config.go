package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Configuration Struct
type Configuration struct {
	Port          int
	GraceShutDown int `json:"serverShutDownGraceDelayInSec"`
	Mongo         struct {
		URI string `json:"URI"`
		DB  string `json:"database"`
	}
}

// Config Struct has config variables
var Config Configuration

// InitiliseConfig func to initilise the config variables dependent on environment
func InitiliseConfig(envStr string) (Configuration, error) {
	filename, err := filepath.Abs("./config/" + envStr + ".json")
	if err != nil {
		return Config, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return Config, err
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return Config, err
	}
	json.Unmarshal(byteValue, &Config)
	return Config, nil
}
