package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"web_complier/configs/application"
	"web_complier/pkg"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	application.AppConfig `yaml:"app"`
	Server                application.ServerConfig `yaml:"server"`
	Mysql                 application.MysqlConfig  `yaml:"mysql"`
	Redis                 application.RedisConfig  `yaml:"redis"`
	Logger                application.LoggerConfig `yaml:"logger"`
}

var Config = &Conf{
	AppConfig: application.App,
	Server:    application.Server,
	Mysql:     application.Mysql,
	Redis:     application.Redis,
	Logger:    application.Logger,
}

func init() {
	// 加载 .yaml 配置
	loadYaml()
}

func loadYaml() {
	currentDirectory, ok := pkg.GetFileDirectoryToCaller()
	if !ok {
		panic("Failed to load configuration: Failed to obtain the current file directory")
	}
	configName := fmt.Sprintf("/configs/config-%s.yaml", os.Getenv("DQENV"))
	fmt.Println("config====:", configName)
	// 生成 config.yaml 文件
	yamlConfig := filepath.Dir(currentDirectory) + configName
	cfg, err := ioutil.ReadFile(yamlConfig)
	if err != nil {
		panic("Failed to read configuration file:" + err.Error())
	}
	err = yaml.Unmarshal(cfg, &Config)
	if err != nil {
		panic("Failed to load configuration:" + err.Error())
	}
}
