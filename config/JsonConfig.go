package config

import (
	"encoding/json"
	"os"
)

func GetJsonConfig(config_path string, e interface{}) error {
	config_file, _ := os.Open(config_path)
	decoder := json.NewDecoder(config_file)
	return decoder.Decode(e)
}
