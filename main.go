package main

import (
	"FuseMiddleware/global"
	"FuseMiddleware/initialize"
)

func main() {
	//初始化路由
	Router := initialize.Routers()

	initialize.InitConfig("./config.yaml")

	global.GVA_DB = initialize.GormMysql()
	defer global.GVA_DB.Close()

	//初始化zap日志
	initialize.InitZap()
	global.GVA_LOG.Info("rest global log")

	Router.Run() // listen and serve on 0.0.0.0:8080
}
