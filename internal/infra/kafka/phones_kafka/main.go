package phones_kafka

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

type PhonesKafka interface {
	UpdatePhone(phone connection.SyncPhoneStruct) error
	CreatePhone(phone connection.SyncPhoneStruct) error
	DeletePhone(phone connection.SyncPhoneStruct) error
}

type phonesKafka struct {
	db *ent.Client
}

func New(db *ent.Client) PhonesKafka {
	return &phonesKafka{
		db: db,
	}
}
