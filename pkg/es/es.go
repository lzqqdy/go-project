package es

import (
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

var Client *elastic.Client

func InitEs() (err error) {
	host := viper.GetString("es.host")
	port := viper.GetString("es.port")
	user := viper.GetString("es.user")
	password := viper.GetString("es.password")
	esUrl := host + ":" + port
	//TODO
	Client, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(esUrl),
		elastic.SetBasicAuth(user, password),
	)
	if err != nil {
		// Handle error
		panic(err)
	}
	defer Client.Stop()
	return
}
