/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-redis licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Mon Jun 26 15:24:19 2023 +0800
 */
package httphandler

import (
	"fmt"
	"net/http"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/redis-plugin/global"
	"openeuler.org/PilotGo/redis-plugin/service"
)

// 安装运行
func InstallRedisExporter(c *gin.Context) {
	// TODOs
	var param common.Batch
	if err := c.BindJSON(&param); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	ret, failRet, err := service.Install(&param)
	if err != nil || failRet != nil {
		response.Fail(c, failRet, err.Error())
		return
	}
	response.Success(c, ret, "安装成功")
}

func UnInstallRedisExporter(c *gin.Context) {
	var param common.Batch
	if err := c.BindJSON(&param); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	ret, failRet, err := service.UnInstall(&param)
	if err != nil || failRet != nil {
		response.Fail(c, failRet, err.Error())
		return
	}
	response.Success(c, ret, "卸载成功")
}

// 重启服务
func RestartRedisExporter(c *gin.Context) {
	var param common.Batch
	if err := c.BindJSON(&param); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	ret, err := service.Restart(&param)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, ret, "重启成功")
}

// 停止服务
func StopRedisExporter(c *gin.Context) {
	var param common.Batch
	if err := c.BindJSON(&param); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	ret, err := service.Stop(&param)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, ret, "停止成功")
}

// 查询数据库安装情况
func GetTargets(c *gin.Context) {
	targets, err := service.GetRedisExporterIp()
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	objs := []RedisexporterObject{
		{
			Targets: targets,
		},
	}
	c.JSON(http.StatusOK, objs)
}

type RedisexporterObject struct {
	Targets []string `json:"targets"`
}

func EnevtMessge(c *gin.Context) {

	err := global.GlobalClient.ListenEvent([]int{1, 2, 3}, []client.EventCallback{func(e *common.EventMessage) {
		fmt.Println("1111111111", e.MessageData)
	}, func(e *common.EventMessage) {
		fmt.Println("2222222222", e.MessageData)
	}, func(e *common.EventMessage) {
		fmt.Println("3333333333", e.MessageData)
	}})

	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "event添加成功")
}

func UNEnevtMessge(c *gin.Context) {
	err := global.GlobalClient.UnListenEvent([]int{1, 2})
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "event删除成功")
}

func PublishEvent(c *gin.Context) {
	msg := common.EventMessage{
		MessageType: 1,
		MessageData: "成功结束",
	}

	err := global.GlobalClient.PublishEvent(msg)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "pubilshevent成功")
}
