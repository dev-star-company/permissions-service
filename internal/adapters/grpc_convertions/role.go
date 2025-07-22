package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
)

func RoleToProto(role *ent.Role) *roles_proto.Role {
	if role == nil {
		return nil
	}

	r := roles_proto.Role{
		Id:          uint32(role.ID),
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt.String(),
	}

	if role.IsActive != nil {
		x := bool(*role.IsActive)
		r.IsActive = &x
	}

	if role.DeletedAt != nil {
		x := role.DeletedAt.Format("2006-01-02 15:04:05")
		r.DeletedAt = &x
	}

	return &r
}
