package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Yaml struct {
	Mysql  mysql  `yaml:"mysql"`
	Http   http   `yaml:"http"`
	Kafka  kafka  `yaml:"kafka"`
	Mongo  mongo  `yaml:"mongo"`
	Thrift thrift `yaml:"thrift"`
}

type mysql struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Dbname       string `yaml:"dbname"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

type http struct {
	Port string `yaml:"port"`
}

type kafka struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type mongo struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type thrift struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func New(filePath string) (*Yaml, error) {
	if file, err := ioutil.ReadFile(filePath); err != nil {
		return nil, err
	} else {
		var conf Yaml
		if err := yaml.Unmarshal(file, &conf); err == nil {
			return &conf, nil
		} else {
			log.Printf("解析出错：%v", err)
			return nil, err
		}
	}
}
