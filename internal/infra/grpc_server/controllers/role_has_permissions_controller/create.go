package role_has_permissions_controller

import (
	"context"
	"permission-service/generated_protos/role_has_permissions_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *role_has_permissions_proto.CreateRequest) (*role_has_permissions_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.RoleHasPermissions.Create().
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &role_has_permissions_proto.CreateResponse{
		RequesterId: uint32(create.CreatedBy),
	}, nil
}
