/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-redis licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Mon Jun 26 15:24:19 2023 +0800
 */
package plugin

import (
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/redis-plugin/config"
)

const Version = "1.0.1"

func Init(plugin *config.PluginRedis, redis *config.RedisServer) *client.PluginInfo {
	PluginInfo := client.PluginInfo{
		Name:        "redis",
		Version:     Version,
		Description: "redis",
		Author:      "wubijie",
		Email:       "wubijie@kylinos.cn",
		Url:         plugin.URL,
		PluginType:  "iframe",
		ReverseDest: "http://" + redis.Addr,
	}
	return &PluginInfo
}

// 请求prometheus插件接口，将gala-ops targets添加到监控清单当中
func addTargets(targets []string, url string) error {
	// TODO:
	// jobName := "redis"
	// url := url+"/api/add_targets"
	/*
	   - job_name: 'redis'
	     static_configs:
	       - targets: ['172.20.32.218:9121']
	*/
	return nil
}

func deleteTargets(targets []string, url string) error {
	// TODO:
	// jobName := "redis"
	// 删除yml文件中有关配置
	return nil
}

func MonitorTargets(targets []string) error {
	plugin, err := client.GetClient().GetPluginInfo("redis")
	if err != nil {
		return err
	}

	if err := addTargets(targets, plugin.Url); err != nil {
		return err
	}

	return nil
}
