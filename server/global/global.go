/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-redis licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: wubijie <wubijie@kylinos.cn>
 * Date: Mon Jun 26 15:24:19 2023 +0800
 */
package global

import (
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
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
