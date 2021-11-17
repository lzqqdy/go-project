package mongodb

import (
	"context"
	"github.com/spf13/viper"
	"go-project/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	Client *mongo.Client
	err    error
)

func InitMongodb() {
	host := viper.GetString("mongodb.host")
	port := viper.GetString("mongodb.port")

	mongodbUrl := host + ":" + port
	// 建立mongodb连接
	clientOptions := options.Client().ApplyURI(mongodbUrl).SetConnectTimeout(5 * time.Second)
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Logger("pkg.mongodb.client").Error(err)
		return
	}
	// 检查连接
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Logger("pkg.mongodb.ping").Error(err)
		return
	}
}
