package kafka_dtos

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

func ToKafkaUser(user ent.User) connection.SyncUserStruct {
	x := connection.SyncUserStruct{
		Uuid:      user.UUID,
		Name:      &user.Name,
		Surname:   &user.Surname,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		CreatedBy: &user.CreatedBy,
		UpdatedBy: &user.UpdatedBy,
		DeletedBy: user.DeletedBy,
	}

	if user.Edges.Phones != nil {
		x.Phones = make([]connection.SyncPhoneStruct, len(user.Edges.Phones))
		for i, phone := range user.Edges.Phones {
			x.Phones[i] = ToKafkaPhone(*phone, user)
		}
	}

	if user.Edges.Emails != nil {
		x.Emails = make([]connection.SyncEmailStruct, len(user.Edges.Emails))
		for i, email := range user.Edges.Emails {
			x.Emails[i] = ToKafkaEmail(*email, user)
		}
	}

	return x
}
