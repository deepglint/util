package config

import (
	"os"
	"text/template"
)

func SetTemplateConfig(template_path string, config_path string, e interface{}) {
	config_file, _ := os.Create(config_path)
	defer config_file.Close()
	t, _ := template.ParseFiles(template_path)
	t.Execute(config_file, e)
}
