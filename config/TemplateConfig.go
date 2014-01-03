package config

import (
	"os"
	"text/template"
)

func SetTemplateConfig(template_path string, config_path string, e interface{}) (err error) {
	config_file, err := os.Create(config_path)
	if err != nil {
		return
	}
	defer config_file.Close()
	t, err := template.ParseFiles(template_path)
	if err != nil {
		return
	}
	return t.Execute(config_file, e)
}
