package kafka_dtos

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

func ToKafkaPhone(phone ent.Phone) connection.SyncPhoneStruct {
	return connection.SyncPhoneStruct{
		Uuid:      phone.UUID,
		Phone:     &phone.Phone,
		CreatedAt: &phone.CreatedAt,
		UpdatedAt: &phone.UpdatedAt,
		DeletedAt: phone.DeletedAt, // This can be nil, so no need for pointer
		CreatedBy: &phone.CreatedBy,
		UpdatedBy: &phone.UpdatedBy,
		DeletedBy: phone.DeletedBy, // This can be nil, so no need for pointer
		Main:      &phone.Main,
	}
}
