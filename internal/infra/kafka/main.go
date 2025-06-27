package kafka

import (
	"permissions-service/internal/app/ent"
	"permissions-service/internal/infra/kafka/emails_kafka"
	"permissions-service/internal/infra/kafka/phones_kafka"
	"permissions-service/internal/infra/kafka/users_kafka"

	"github.com/dev-star-company/kafka-go/connection"
)

type kafka struct {
	Db *ent.Client
	c  *connection.Connectioner

	UsersKafka  users_kafka.UsersKafka
	EmailsKafka emails_kafka.EmailsKafka
	PhonesKafka phones_kafka.PhonesKafka
}

func New(Db *ent.Client, conn *connection.Connectioner) *kafka {
	return &kafka{
		Db:          Db,
		c:           conn,
		UsersKafka:  users_kafka.New(Db),
		EmailsKafka: emails_kafka.New(Db),
		PhonesKafka: phones_kafka.New(Db),
	}
}
