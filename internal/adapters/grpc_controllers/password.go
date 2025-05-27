package grpc_controllers

import (
	"permission-service/generated_protos/users_proto"
	"permission-service/internal/app/ent"
)

func PasswordToProto(password *ent.Password) *users_proto.Password {
	if password == nil {
		return nil
	}
	p := users_proto.Password{
		Id:        uint32(password.ID),
		Password:  password.Password,
		CreatedAt: password.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: password.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy: uint32(password.CreatedBy),
		UpdatedBy: uint32(password.UpdatedBy),
	}

	if password.DeletedAt != nil {
		x := password.DeletedAt.Format("2006-01-02 15:04:05")
		p.DeletedAt = &x
	}

	if password.DeletedBy != nil {
		x := uint32(*password.DeletedBy)
		p.DeletedBy = &x
	}

	return &p
}
