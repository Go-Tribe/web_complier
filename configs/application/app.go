package application

import (
	"web_complier/pkg"
	"web_complier/pkg/convert"
)

type AppConfig struct {
	AppEnv         string `yaml:"app_env"`
	Debug          bool   `yaml:"debug"`
	Language       string `yaml:"language"`
	StaticBasePath string `yaml:"base_path"`
}

var App = AppConfig{
	AppEnv:         "local",
	Debug:          true,
	Language:       "zh_CN",
	StaticBasePath: getDefaultPath(),
}

func getDefaultPath() (path string) {
	path = pkg.GetRunPath()
	path, _ = convert.GetString(pkg.If(path != "", path, "/tmp"))
	return
}
