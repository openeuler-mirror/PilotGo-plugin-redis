package config

import (
	"flag"
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	con "gitee.com/openeuler/PilotGo/sdk/utils/config"

	"gopkg.in/yaml.v2"
)

type PluginRedis struct {
	URL        string `yaml:"url"`
	PluginType string `yaml:"plugin_type"`
}

type RedisServer struct {
	Addr string `yaml:"addr"`
}

type PilotGoServer struct {
	Addr string `yaml:"addr"`
}

type MysqlDBInfo struct {
	HostName string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"user"`
	Password string `yaml:"password"`
	DataBase string `yaml:"database"`
}

type ServerConfig struct {
	PluginRedis   *PluginRedis    `yaml:"plugin_redis"`
	RedisServer   *RedisServer    `yaml:"redis_server"`
	PilotGoServer *PilotGoServer  `yaml:"pilotgo_server"`
	Logopts       *logger.LogOpts `yaml:"log"`
	Mysql         *MysqlDBInfo    `yaml:"mysql"`
}

var global_config ServerConfig
var Config_file string

func Init() error {
	flag.StringVar(&Config_file, "conf", "./config.yaml", "plugin-resid configuration file")
	flag.Parse()
	return con.Load(Config_file, &global_config)
}

func Config() *ServerConfig {
	return &global_config
}

func readConfig(file string, config interface{}) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("open %s failed! err = %s\n", file, err.Error())
		return err
	}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		fmt.Printf("yaml Unmarshal %s failed!\n", string(bytes))
		return err
	}
	return nil
}
