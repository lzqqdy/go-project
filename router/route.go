package router

import (
	"go-project/app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
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
	// Websocket服务，路由ws://localhost:8000/ws/xxx，xxx对应Handlers中方法定位到不同路由
	g.GET("/ws/:business", controller.Websocket)
	//中间件Token验证
	g.Use(middleware.Token())
	{
		g.GET("/test", controller.GetTest)
	}
	//...
	return g
}
