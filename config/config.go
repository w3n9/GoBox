package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type YamlConfig struct{
	Project struct{
		Name string `yaml:"name"`
	} `yaml:"project"`
	Storage struct{
		Path string `yaml:"path"`
	} `yaml:"storage"`
	File struct{
		Encrypt string `yaml:"encrypt"`
	} `yaml:"file"`
	Web struct{
		StaticUrlPrefix string `yaml:"staticUrlPrefix"`
		StaticDirPrefix string `yaml:"staticDirPrefix"`
	} `yaml:"web"`
	DataSource struct{
		Type string `yaml:"type"`
		Url string `yaml:"url"`
	} `yaml:"dataSource"`
}
var conf YamlConfig

func init(){
	yml, err := os.Open("./conf.yml")
	if err != nil {
		panic(err)
	}
	decoder := yaml.NewDecoder(yml)
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
}
func Get() YamlConfig{
	return conf
}