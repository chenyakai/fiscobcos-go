package core

import (
	"fmt"
	"github.com/chenyakai/fiscobcos-go/config"
	"github.com/chenyakai/fiscobcos-go/whole"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// InitConf 读取yaml文件配置
func InitConf(path string) {
	//const ConfigFile = "resources/settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")
	whole.Config = c
}
