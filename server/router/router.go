package router

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/redis-plugin/global"
	"openeuler.org/PilotGo/redis-plugin/httphandler"
)

// gin.egnine充当server的角色
func InitRouter() *gin.Engine {
	//确定程序如何处理请求，默认使用标准http和基于url的模式路由
	//使用此模式会禁用活限制某些权限，例如自动跟新代码和模拟响应
	gin.SetMode(gin.ReleaseMode)
	//创建一个新的http服务器实例，可以添加中间件、路由、处理函数等组件，构建自己的程序
	router := gin.New()
	//添加日志中间件
	router.Use(logger.LoggerDebug())
	//添加恢复中间件
	router.Use(gin.Recovery())
	return router
}

func RegisterAPIs(router *gin.Engine) {
	//输出插件初始化的信息
	global.GlobalClient.RegisterHandlers(router)

	pg := router.Group("/plugin/" + global.GlobalClient.PluginInfo.Name)
	{
		pg.POST("/install", httphandler.InstallRedisExporter)
		pg.POST("/remove", httphandler.UnInstallRedisExporter)
		pg.POST("/restart", httphandler.RestartRedisExporter)
		pg.POST("/stop", httphandler.StopRedisExporter)
		pg.GET("/targets", httphandler.GetTargets)
	}
}
