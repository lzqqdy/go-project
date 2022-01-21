package mq

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/rs/xid"
)

func InitConsumer(redisClient *redis.Client, subject string, consumersGroup string) (err error) {
	err = redisClient.XGroupCreate(subject, consumersGroup, "0").Err()
	if err != nil {
		log.Println(err)
		return
	}

	uniqueID := xid.New().String()
	go func() {
		for {
			entries, err := redisClient.XReadGroup(&redis.XReadGroupArgs{
				Group:    consumersGroup,
				Consumer: uniqueID,
				Streams:  []string{subject, ">"},
				Count:    2,
				Block:    0,
				NoAck:    false,
			}).Result()
			if err != nil {
				log.Fatal(err)
				return
			}

			for i := 0; i < len(entries[0].Messages); i++ {
				messageID := entries[0].Messages[i].ID
				values := entries[0].Messages[i].Values
				eventDescription := fmt.Sprintf("%v", values["whatHappened"])
				ticketID := fmt.Sprintf("%v", values["ticketID"])
				ticketData := fmt.Sprintf("%v", values["ticketData"])

				if eventDescription == "ticket received" {
					err := handleNewTicket(ticketID, ticketData)
					if err != nil {
						log.Fatal(err)
						return
					}
					redisClient.XAck(subject, consumersGroup, messageID)
				}
			}
		}
	}()
	return
}
func handleNewTicket(ticketID string, ticketData string) error {
	log.Printf("Handling new ticket id : %s data %s\n", ticketID, ticketData)
	//TODO

	return nil
}
