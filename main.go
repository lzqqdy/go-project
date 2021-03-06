package main

import (
	"go-project/crontab"
	"go-project/models"
	"go-project/pkg/config"
	"go-project/pkg/es"
	"go-project/pkg/logger"
	"go-project/pkg/mongodb"
	"go-project/pkg/mq"
	"go-project/pkg/redis"
	"go-project/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/spf13/pflag"
)

func init() {
	var conf = pflag.StringP("config", "c", "", "config filepath")

	pflag.Parse()
	config.InitConfig(*conf)
	models.InitMysql()
	redis.InitRedis()
	logger.Initiate()
	es.InitEs()
	mongodb.InitMongodb()
}

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name go-project
// @BasePath /
func main() {
	//定时任务
	crontab.Timer()
	//消息队列
	mq.InitMq()
	// gin
	gin.SetMode(viper.GetString("app.mode"))
	r := gin.New()
	r = router.Load(r)
	err := r.Run(viper.GetString("app.addr"))
	if err != nil {
		panic(err)
	}
}
