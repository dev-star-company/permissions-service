package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"
)

func PermissionToProto(permission *ent.Permission) *permission_proto.Permission {
	if permission == nil {
		return nil
	}

	p := permission_proto.Permission{
		Id:          uint32(permission.ID),
		Name:        permission.Name,
		Description: permission.Description,
		IsActive:    permission.IsActive,
	}

	return &p
}
