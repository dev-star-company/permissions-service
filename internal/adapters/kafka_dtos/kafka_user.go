package kafka_dtos

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

func ToKafkaUser(user ent.User) connection.SyncUserStruct {
	return connection.SyncUserStruct{
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
}
