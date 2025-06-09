package grpc_controllers

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"
)

func RoleHasPermissionsToProto(role_has_permissions *ent.RoleHasPermissions) *role_has_permissions_proto.RoleHasPermissions {
	if role_has_permissions == nil {
		return nil
	}

	e := role_has_permissions_proto.RoleHasPermissions{
		RoleId:       uint32(role_has_permissions.RoleID),
		PermissionId: uint32(role_has_permissions.PermissionID),
		CreatedAt:    role_has_permissions.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    role_has_permissions.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy:    uint32(role_has_permissions.CreatedBy),
		UpdatedBy:    uint32(role_has_permissions.UpdatedBy),
	}

	if role_has_permissions.DeletedAt != nil {
		x := role_has_permissions.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	if role_has_permissions.DeletedBy != nil {
		x := uint32(*role_has_permissions.DeletedBy)
		e.DeletedBy = &x
	}

	return &e
}
