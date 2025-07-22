package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
)

func PasswordToProto(password *ent.Password) *auth_users_proto.Password {
	if password == nil {
		return nil
	}
	p := auth_users_proto.Password{
		Id:        uint32(password.ID),
		Password:  password.Password,
		CreatedAt: password.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if password.DeletedAt != nil {
		x := password.DeletedAt.Format("2006-01-02 15:04:05")
		p.DeletedAt = &x
	}

	return &p
}
