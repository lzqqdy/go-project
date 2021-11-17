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

func Index(g *gin.Context) {
	api.Success("SUCCESS").End(g)
}
func Ping(g *gin.Context) {
	//日志测试
	logger.Logger("app.index.ping").Error("pong")
	//redis测试
	key, _ := redis.Client.Get("test").Result()
	fmt.Println(key)
	//ES测试
	logic.TestEs()
	//MongoDB测试
	logic.TestMongoDB()
	api.Error("pong").End(g)
}

//Test 参数验证
type Test struct {
	Name string `form:"name" binding:"required"`
}

func GetTest(g *gin.Context) {
	name := g.Query("name")
	var form Test
	if err := g.ShouldBind(&form); err != nil {
		api.Error(fmt.Sprint(err)).End(g)
		return
	}
	state := -1
	testLogic := logic.Test{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(g),
		PageSize: 10,
	}

	test, err := testLogic.GetAll()
	if err != nil {
		api.Error("error").End(g)
		return
	}

	count, err := testLogic.Count()
	if err != nil {
		api.Error("error").End(g)
		return
	}
	api.Success(map[string]interface{}{
		"lists": test,
		"total": count,
	}).End(g)
}
