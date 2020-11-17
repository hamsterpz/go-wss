package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func YamlConfReader(path string, out interface{}) interface{} {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("文件读取失败: %v", err)
	}
	err = yaml.Unmarshal(data, out)
	if err != nil {
		log.Panicf("序列化失败失败: %v", err)
	}
	return out
}
