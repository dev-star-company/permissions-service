package grpc_controllers

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/users_proto"
)

func EmailToProto(email *ent.Email) *users_proto.Email {
	if email == nil {
		return nil
	}

	e := users_proto.Email{
		Id:        uint32(email.ID),
		Email:     email.Email,
		CreatedAt: email.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: email.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy: uint32(email.CreatedBy),
		UpdatedBy: uint32(email.UpdatedBy),
		Main:      email.Main,
	}

	if email.DeletedAt != nil {
		x := email.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	if email.DeletedBy != nil {
		x := uint32(*email.DeletedBy)
		e.DeletedBy = &x
	}

	return &e
}
