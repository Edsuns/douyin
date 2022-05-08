package config

import (
	"douyin/pkg/database"
	"douyin/pkg/security"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

type Config struct {
	Port string

	Gorm struct {
		LogLevel string `yaml:"log-level"`
	}
	Mysql database.MysqlConfig

	Jwt security.JwtConfig
}

var Val *Config

func Setup(config string) {
	var err error
	Val, err = loadConfig(config)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	checkRequired(*config)

	return config, nil
}

func checkRequired(st interface{}) {
	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)
	for k := 0; k < t.NumField(); k++ {
		fieldType := v.Field(k).Kind()
		if fieldType == reflect.Struct {
			checkRequired(v.Field(k).Interface())
		}
		if t.Field(k).Tag.Get("config") != "optional" {
			if v.Field(k).IsZero() {
				panic(fmt.Sprintf("configuration %+v can not be zero", t.Field(k).Name))
			}
		}
	}
}
