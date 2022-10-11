package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mygin.com/mygin/common"
	"mygin.com/mygin/middleware"
)

func main() {
	// 初始化配置
	common.InitConfig()
	// 建立数据库连接
	db := common.InitDB()
	defer db.Close()
	// 获取路由引擎
	engine := gin.Default()
	// 配置跨域访问
	engine.Use(middleware.CorsMiddleware())

	// 路由配置
	CollectRoute(engine)

	var port string = viper.GetString("server.port")
	if err := engine.Run(":" + port); err != nil {
		panic("server is shutdown err: " + err.Error())
	}
}
