package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
)

func PhoneToProto(phone *ent.Phone) *auth_users_proto.Phone {
	if phone == nil {
		return nil
	}
	p := auth_users_proto.Phone{
		Id:        uint32(phone.ID),
		Phone:     phone.Phone,
		CreatedBy: uint32(phone.CreatedBy),
		UpdatedBy: uint32(phone.UpdatedBy),
		CreatedAt: phone.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: phone.UpdatedAt.Format("2006-01-02 15:04:05"),
		Main:      phone.Main,
	}

	if phone.DeletedAt != nil {
		x := phone.DeletedAt.Format("2006-01-02 15:04:05")
		p.DeletedAt = &x
	}

	if phone.DeletedBy != nil {
		x := uint32(*phone.DeletedBy)
		p.DeletedBy = &x
	}

	return &p
}
