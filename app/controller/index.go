package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/app/logic"
	"go-project/pkg/api"
	"go-project/pkg/logger"
	"go-project/pkg/redis"
	"go-project/pkg/util"
)

func Index(c *gin.Context) {
	api.Success("SUCCESS").End(c)
}
func Ping(c *gin.Context) {
	//日志测试
	logger.Logger("app.index.ping").Error("pong")
	//redis连接测试
	key, _ := redis.Client.Get("test").Result()
	fmt.Println(key)

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
