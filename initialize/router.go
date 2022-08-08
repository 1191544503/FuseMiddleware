package initialize

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// var RouterGroupApp = new(m.RouterGroup)

	// systemRouter := RouterGroupApp.Util

	// publicGroup := Router.Group("")
	// systemRouter.InitUtilRouter(publicGroup)

	return Router
}
