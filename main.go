package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-project/crontab"
	"go-project/models"
	"go-project/pkg/config"
	"go-project/pkg/logger"
	"go-project/pkg/redis"
	"go-project/router"

	"github.com/spf13/pflag"
)

func init() {
	var conf = pflag.StringP("config", "c", "", "config filepath")

	pflag.Parse()
	config.InitConfig(*conf)
	models.InitMysql()
	redis.InitRedis()
	logger.Initiate()
}

func main() {

	//定时任务
	crontab.Timer()

	// gin
	gin.SetMode(viper.GetString("app.mode"))
	r := gin.New()
	r = router.Load(r)
	err := r.Run(viper.GetString("app.addr"))
	if err != nil {
		panic(err)
	}
}
