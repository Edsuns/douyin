package config

import (
	"douyin/pkg/dbx"
	"douyin/pkg/security"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

type Config struct {
	Port string

	Static struct {
		BaseUrl  string `yaml:"base-url"`
		Route    string
		Filepath string
	}

	Gorm struct {
		LogLevel string `yaml:"log-level"`
	}
	Mysql dbx.MysqlConfig

	Jwt security.JwtConfig
}

var Val *Config

func Load(path string) {
	var err error
	Val, err = loadConfig(filepath.Join(path, "config.yaml"))
	if err != nil {
		panic(err)
	}

	// resolve relative paths
	Val.Static.Filepath = filepath.Join(path, Val.Static.Filepath)
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

	// override config if env variables exist
	loadConfigFromEnv(config)

	checkRequired(*config)

	return config, nil
}

func loadConfigFromEnv(config *Config) {
	// static-url-base
	if baseUrl := os.Getenv("static-base-url"); baseUrl != "" {
		config.Static.BaseUrl = baseUrl
	}
	// db-name
	if dbName := os.Getenv("db-name"); dbName != "" {
		config.Mysql.Database = dbName
	}
	// db-host
	if host := os.Getenv("db-host"); host != "" {
		config.Mysql.Host = host
	}
	// db-username
	if username := os.Getenv("db-username"); username != "" {
		config.Mysql.Username = username
	}
	// db-password
	if password := os.Getenv("db-password"); password != "" {
		config.Mysql.Password = password
	}
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
