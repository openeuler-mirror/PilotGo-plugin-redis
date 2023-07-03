package global

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
	"gorm.io/gorm"
)

var (
	GlobalClient *client.Client
	GlobalDB     *gorm.DB
)

// 执行操作状态
const (
	StatusInstall = "安装"
	StatusRemove  = "卸载"
	StatusErroe   = "安装错误"
)
