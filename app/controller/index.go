package controller

import (
	"github.com/gin-gonic/gin"
	"go-project/app/logic"
	"go-project/pkg/api"
	"go-project/pkg/logger"
	"go-project/pkg/util"
)

func Index(c *gin.Context) {
	api.Success("SUCCESS").End(c)
}
func Ping(c *gin.Context) {
	logger.Logger("app.index.ping").Error("pong")
	api.Error("pong").End(c)
}
func GetTest(c *gin.Context) {
	name := c.Query("name")
	state := -1
	testLogic := logic.Test{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: 10,
	}

	test, err := testLogic.GetAll()
	if err != nil {
		api.Error("error").End(c)
		return
	}

	count, err := testLogic.Count()
	if err != nil {
		api.Error("error").End(c)
		return
	}
	api.Success(map[string]interface{}{
		"lists": test,
		"total": count,
	}).End(c)
}
