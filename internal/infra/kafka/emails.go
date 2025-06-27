package kafka

import (
	"context"
	"fmt"

	"github.com/dev-star-company/kafka-go/actions"
)

func (k *kafka) SyncEmails() {
	incomingEmails, err := k.c.SubscribeToEmails(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		r := <-incomingEmails
		switch msg := r.Message; msg.Action {
		case actions.CREATE:
			err := k.EmailsKafka.CreateEmail(msg.Payload)
			if err != nil {
				fmt.Println("error creating email: " + err.Error())
			} else {
				r.CommitFn()
			}
		case actions.UPDATE:
			err := k.EmailsKafka.UpdateEmail(msg.Payload)
			if err != nil {
				fmt.Println("error updating user: " + err.Error())
			} else {
				r.CommitFn()
			}
		case actions.DELETE:
			err := k.EmailsKafka.DeleteEmail(msg.Payload)
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
