package kafka

import (
	"context"
	"fmt"

	"github.com/dev-star-company/kafka-go/actions"
)

func (k *kafka) SyncPhones() {
	incomingPhones, err := k.c.SubscribeToPhones(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		r := <-incomingPhones
		switch msg := r.Message; msg.Action {
		case actions.CREATE:
			err := k.PhonesKafka.CreatePhone(msg.Payload)
			if err != nil {
				fmt.Println("error creating phone: " + err.Error())
			} else {
				r.CommitFn()
			}
		case actions.UPDATE:
			err := k.PhonesKafka.UpdatePhone(msg.Payload)
			if err != nil {
				fmt.Println("error updating phone: " + err.Error())
			} else {
				r.CommitFn()
			}
		case actions.DELETE:
			err := k.PhonesKafka.DeletePhone(msg.Payload)
			if err != nil {
				fmt.Println("error deleting phone: " + err.Error())
			} else {
				r.CommitFn()
			}
		default:
			fmt.Println("unknown action: " + msg.Action)
		}
	}
}
