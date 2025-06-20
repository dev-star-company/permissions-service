package kafka_dtos

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

func ToKafkaEmail(email ent.Email) connection.SyncEmailStruct {
	return connection.SyncEmailStruct{
		ID:        email.ID,
		Email:     &email.Email,
		CreatedAt: &email.CreatedAt,
		UpdatedAt: &email.UpdatedAt,
		DeletedAt: email.DeletedAt, // This can be nil, so no need for pointer
		CreatedBy: &email.CreatedBy,
		UpdatedBy: &email.UpdatedBy,
		DeletedBy: email.DeletedBy, // This can be nil, so no need for pointer
		Main:      &email.Main,
	}
}
