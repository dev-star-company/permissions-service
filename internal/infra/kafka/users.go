package kafka

import (
	"context"
	"fmt"
	"permissions-service/internal/config/env"

	"github.com/dev-star-company/kafka-go/actions"
	"github.com/dev-star-company/kafka-go/topics"
)

func (k *kafka) SyncUsers() {
	_, err := k.c.ConnectToTopic(topics.SyncUsers, env.KAFKA_BROKER_URL)
	if err != nil {
		fmt.Println("error connecting to topic SyncUsers: " + err.Error())
	}

	incomingUsers, err := k.c.SubscribeToUsers(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		switch msg := <-incomingUsers; msg.Action {
		case actions.CREATE:
			err := k.UsersKafka.CreateUser(msg.Payload)
			if err != nil {
				fmt.Println("error creating user: " + err.Error())
			}
		case actions.UPDATE:
			err := k.UsersKafka.UpdateUser(msg.Payload)
			if err != nil {
				fmt.Println("error updating user: " + err.Error())
			}
		case actions.DELETE:
			err := k.UsersKafka.DeleteUser(msg.Payload)
			if err != nil {
				fmt.Println("error deleting user: " + err.Error())
			}
		default:
			fmt.Println("unknown action: " + msg.Action)
		}
	}
}
