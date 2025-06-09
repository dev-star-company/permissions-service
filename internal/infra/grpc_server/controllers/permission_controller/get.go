package permission_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/permission"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *permission_proto.GetRequest) (*permission_proto.GetResponse, error) {
	permission, err := c.Db.Permission.
		Query().
		Where(permission.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.PermissionNotFound(int(in.Id))
	}

	return &permission_proto.GetResponse{
		Name:         permission.Name,
		Description:  permission.Description,
		InternalName: permission.InternalName,
		IsActive:     permission.IsActive,
		ServiceId:    uint32(permission.ServiceID),
	}, nil
}
