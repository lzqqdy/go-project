package mq

import (
	"go-project/pkg/redis"
	"log"
	"math/rand"

	"github.com/spf13/viper"
)

func InitMq() {
	status := viper.GetBool("mq.status")
	if status {
		subject := viper.GetString("mq.subject")
		consumersGroup := viper.GetString("mq.consumersGroup")
		value := map[string]interface{}{
			"whatHappened": string("ticket received"),
			"ticketID":     int(rand.Intn(100000000)),
			"ticketData":   string("some ticket data"),
		}
		err := PublishTicketReceivedEvent(redis.Client, subject, value)
		if err != nil {
			log.Fatal(err)
		}
		err = InitConsumer(redis.Client, subject, consumersGroup)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
