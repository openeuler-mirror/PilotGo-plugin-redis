package global

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
)

var (
	GlobalClient *client.Client
)

// 执行操作状态
const (
	StatusInstall = "安装"
	StatusRemove  = "卸载"
	StatusError   = "执行错误"
)
