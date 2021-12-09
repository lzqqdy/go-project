package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-project/app/middleware"

	"go-project/app/controller"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-project/docs"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})
	g.GET("/", controller.Index)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.GET("/ping", controller.Ping)

	g.GET("/getTest", controller.GetTest)
	//中间件Token验证
	g.Use(middleware.Token())
	{
		g.GET("/test", controller.GetTest)
	}
	//...
	return g
}
