package httphandler

import (
	"net/http"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
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
