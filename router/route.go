package router

import (
	"net/http"

	"go-project/app/controller"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	g.GET("/", controller.Index)
	g.GET("/ping", controller.Ping)
	//...
	g.GET("/getTest", controller.GetTest)

	return g
}
