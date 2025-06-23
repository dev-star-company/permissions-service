package kafka_dtos

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

func ToKafkaPhone(phone ent.Phone, user ent.User) connection.SyncPhoneStruct {
	return connection.SyncPhoneStruct{
		Uuid:      phone.UUID,
		UserUuid:  user.UUID,
		Phone:     &phone.Phone,
		CreatedAt: &phone.CreatedAt,
		UpdatedAt: &phone.UpdatedAt,
		DeletedAt: phone.DeletedAt,
		CreatedBy: &phone.CreatedBy,
		UpdatedBy: &phone.UpdatedBy,
		DeletedBy: phone.DeletedBy,
		Main:      &phone.Main,
	}
}
