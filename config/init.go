package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func CfgInit() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	//注册 config 配置文件
	var conf Cfg
	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		log.Fatalf("无法解析 YAML 文件: %v", err)
	}

	SetGbl(conf)
}
