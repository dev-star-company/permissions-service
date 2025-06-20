package users_kafka

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

type UsersKafka interface {
	UpdateUser(user connection.SyncUserStruct) error
	CreateUser(user connection.SyncUserStruct) error
	DeleteUser(user connection.SyncUserStruct) error
}

type usersKafka struct {
	db *ent.Client
}

func New(db *ent.Client) UsersKafka {
	return &usersKafka{
		db: db,
	}
}
