package role_has_permissions_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/rolehaspermissions"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *role_has_permissions_proto.GetRequest) (*role_has_permissions_proto.GetResponse, error) {
	_, err := c.Db.RoleHasPermissions.
		Query().
		Where(rolehaspermissions.RoleID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.RoleHasPermissionNotFound(int(in.Id))
	}

	return &role_has_permissions_proto.GetResponse{
		RequesterUuid: in.RequesterUuid,
	}, nil
}
