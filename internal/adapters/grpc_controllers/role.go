package grpc_controllers

import (
	"permission-service/generated_protos/roles_proto"
	"permission-service/internal/app/ent"
)

func RoleToProto(role *ent.Role) *roles_proto.Role {
	r := roles_proto.Role{
		Id:          uint32(role.ID),
		Name:        role.Name,
		Description: role.Description,
		IsActive:    role.IsActive,
		CreatedAt:   role.CreatedAt.String(),
		UpdatedAt:   role.UpdatedAt.String(),
		CreatedBy:   uint32(role.CreatedBy),
		UpdatedBy:   uint32(role.UpdatedBy),
	}

	if role.DeletedAt != nil {
		x := role.DeletedAt.Format("2006-01-02 15:04:05")
		r.DeletedAt = &x
	}

	if role.DeletedBy != nil {
		x := uint32(*role.DeletedBy)
		r.DeletedBy = &x
	}

	return &r
}
