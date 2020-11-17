package config

import "go-wss/src/util"

type systemConf struct {
	Port  int `yaml:"SERVER_PORT"`
	Mongo struct {
		HOST     string `yaml:"HOST"`
		USERNAME string `yaml:"USERNAME"`
		PASSWORD string `yaml:"PASSWORD"`
		DATABASE string `yaml:"DATABASE"`
		PORT     int    `yaml:"PORT"`
	} `yaml:"MONGO"`
}

var System = util.YamlConfReader("config.yaml", &systemConf{}).(*systemConf)
