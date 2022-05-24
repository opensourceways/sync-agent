package core

import (
	"github.com/gin-gonic/gin"
	"github.com/opensourceways/sync-agent/config"
	"github.com/opensourceways/sync-agent/controller/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/opensourceways/sync-agent/docs"
	"github.com/opensourceways/sync-agent/middleware"
	"github.com/opensourceways/sync-agent/router"
	"github.com/opensourceways/sync-agent/utils"
)

func initRouter() *gin.Engine {
	if config.Config().GetEnv() == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	gRouter := gin.New()

	gRouter.Use(middleware.Logger(), gin.Recovery())

	gRouter.NoRoute(func(c *gin.Context) {
		utils.NotFoundError(c)
	})

	//setting swagger doc address
	gRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//router group
	apiGroup := gRouter.Group("/v1")

	// register synchronization router group
	router.NewRouteRegister(router.RouteGP(apiGroup), router.RouteModules(&v1.SyncController{})).Do()

	return gRouter
}
