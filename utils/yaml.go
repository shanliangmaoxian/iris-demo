package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configFile []byte

type Yaml struct {
	Server struct {
		Port    string `yaml:"port"`
		Charset string `yaml:"charset"`
	}
	Mysql struct {
		Port   string `yaml:"port"`
		Host   string `yaml:"host"`
		User   string `yaml:"user"`
		Pwd    string `yaml:"pwd"`
		Dbname string `yaml:"dbname"`
	}
}

func GetYaml() (e *Yaml, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
}
