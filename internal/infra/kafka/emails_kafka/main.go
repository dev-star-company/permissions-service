package emails_kafka

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

type EmailsKafka interface {
	UpdateEmail(email connection.SyncEmailStruct) error
	CreateEmail(email connection.SyncEmailStruct) error
	DeleteEmail(email connection.SyncEmailStruct) error
}

type emailsKafka struct {
	db *ent.Client
}

func New(db *ent.Client) EmailsKafka {
	return &emailsKafka{
		db: db,
	}
}
