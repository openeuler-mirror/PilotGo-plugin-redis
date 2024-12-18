/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-redis licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Mon Jun 26 15:24:19 2023 +0800
 */
package main

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/redis-plugin/config"
	"openeuler.org/PilotGo/redis-plugin/db"
	"openeuler.org/PilotGo/redis-plugin/global"
	"openeuler.org/PilotGo/redis-plugin/plugin"
	"openeuler.org/PilotGo/redis-plugin/router"
)

/*
-redis.addr：指明 Redis 节点的地址，默认为 redis://localhost:6379(如果有多个redis实例, redis_exporter作者建议启动多个redis_exporter进程来进行监控数据获取)
-redis.password：验证 Redis 时使用的密码；
-redis.file：包含一个或多个redis 节点的文件路径，每行一个节点，此选项与 -redis.addr 互斥。
-web.listen-address：监听的地址和端口，默认为 0.0.0.0:9121
*/
func main() {
	fmt.Println("hello redis")
	err := config.Init()
	if err != nil {
		fmt.Println("failed to load configure, exit..", err)
		os.Exit(-1)
	}

	if err := logger.Init(config.Config().Logopts); err != nil {
		logger.Error("logger init failed, please check the config file: %s", err)
		os.Exit(-1)
	}

	if err := db.MysqldbInit(config.Config().Mysql); err != nil {
		logger.Error("mysql db init failed, please check again: %s", err)
		os.Exit(-1)
	}

	server := router.InitRouter()
	global.GlobalClient = client.DefaultClient(plugin.Init(config.Config().PluginRedis, config.Config().RedisServer))

	go router.RegisterAPIs(server)
	if err := server.Run(config.Config().HttpServer.Addr); err != nil {
		logger.Error("failed to run server: %s", err)
		os.Exit(-1)
	}
}
