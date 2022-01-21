package mq

import (
	"log"

	"github.com/go-redis/redis"
)

func PublishTicketReceivedEvent(client *redis.Client, subject string, streamValue map[string]interface{}) error {
	log.Println("Publishing event to Redis")
	err := client.XAdd(&redis.XAddArgs{
		Stream:       subject,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       streamValue,
	}).Err()
	return err
}
