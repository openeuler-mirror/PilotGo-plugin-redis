/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-redis licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Mon Jun 26 15:24:19 2023 +0800
 */
package config

import (
	"flag"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	con "gitee.com/openeuler/PilotGo/sdk/utils/config"
)

type PluginRedis struct {
	URL string `yaml:"url"`
}

type RedisServer struct {
	Addr string `yaml:"addr"`
}

type HttpServer struct {
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
	PluginRedis *PluginRedis    `yaml:"plugin_redis"`
	RedisServer *RedisServer    `yaml:"redis_server"`
	HttpServer  *HttpServer     `yaml:"http_server"`
	Logopts     *logger.LogOpts `yaml:"log"`
	Mysql       *MysqlDBInfo    `yaml:"mysql"`
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
