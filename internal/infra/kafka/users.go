package kafka

import (
	"context"
	"fmt"
	"permissions-service/internal/config/env"

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
		case "create":
		}
	}
}
