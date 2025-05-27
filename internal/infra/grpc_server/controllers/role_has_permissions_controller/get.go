package role_has_permissions_controller

import (
	"context"
	"permission-service/generated_protos/role_has_permissions_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/app/ent/rolehaspermissions"
	"permission-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *role_has_permissions_proto.GetRequest) (*role_has_permissions_proto.GetResponse, error) {
	role_has_permissions, err := c.Db.RoleHasPermissions.
		Query().
		Where(rolehaspermissions.RoleID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.RoleHasPermissionNotFound(int(in.Id))
	}

	return &role_has_permissions_proto.GetResponse{
		RequesterId: uint32(role_has_permissions.CreatedBy),
	}, nil
}
