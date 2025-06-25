package kafka

import (
	"context"
	"fmt"

	"github.com/dev-star-company/kafka-go/actions"
)

func (k *kafka) SyncUsers() {
	incomingUsers, err := k.c.SubscribeToUsers(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		r := <-incomingUsers
		switch msg := r.Message; msg.Action {
		case actions.CREATE:
			err := k.UsersKafka.CreateUser(msg.Payload)
			if err != nil {
				fmt.Println("error creating user: " + err.Error())
			} else {
				r.CommitFn()
			}
		case actions.UPDATE:
			err := k.UsersKafka.UpdateUser(msg.Payload)
			if err != nil {
				fmt.Println("error updating user: " + err.Error())
			} else {
				r.CommitFn()
			}
		case actions.DELETE:
			err := k.UsersKafka.DeleteUser(msg.Payload)
			if err != nil {
				fmt.Println("error deleting user: " + err.Error())
			} else {
				r.CommitFn()
			}
		default:
			fmt.Println("unknown action: " + msg.Action)
		}
	}
}
