package json

import (
	"encoding/json"
	"os"
	"strings"
)

func DeserializeFromFile(config_path string, e interface{}) (err error) {
	config_file, _ := os.Open(config_path)
	if err != nil {
		return
	}
	decoder := json.NewDecoder(config_file)
	return decoder.Decode(e)
}

func DeserializeFromString(json_string string, e interface{}) (err error) {
	decoder := json.NewDecoder(strings.NewReader(json_string))
	return decoder.Decode(e)
}
