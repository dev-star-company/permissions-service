package kafka_dtos

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

func ToKafkaUser(user ent.User) connection.SyncUserStruct {
	return connection.SyncUserStruct{
		ID:        user.ID,
		Name:      &user.Name,
		Surname:   &user.Surname,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		DeletedAt: user.DeletedAt, // This can be nil, so no need for pointer
		CreatedBy: &user.CreatedBy,
		UpdatedBy: &user.UpdatedBy,
		DeletedBy: user.DeletedBy, // This can be nil, so no need for pointer
	}
}
