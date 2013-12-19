package config

import (
	"encoding/json"
	"os"
	"strings"
)

func GetJsonConfigFromFile(config_path string, e interface{}) error {
	config_file, _ := os.Open(config_path)
	decoder := json.NewDecoder(config_file)
	return decoder.Decode(e)
}

func GetJsonConfigFromString(json_string string, e interface{}) error {
	decoder := json.NewDecoder(strings.NewReader(json_string))
	return decoder.Decode(e)
}
