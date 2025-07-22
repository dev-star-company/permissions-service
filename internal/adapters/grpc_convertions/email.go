package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
)

func EmailToProto(email *ent.Email) *auth_users_proto.Email {
	if email == nil {
		return nil
	}

	e := auth_users_proto.Email{
		Id:        uint32(email.ID),
		Email:     email.Email,
		CreatedAt: email.CreatedAt.Format("2006-01-02 15:04:05"),
		Main:      email.Main,
	}

	if email.DeletedAt != nil {
		x := email.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	return &e
}
