package role_has_permissions_controller

import (
	"context"

	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *role_has_permissions_proto.CreateRequest) (*role_has_permissions_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	_, err = c.Db.RoleHasPermissions.Create().
		SetRoleID(int(in.RoleId)).
		SetPermissionID(int(in.PermissionId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("role_has_permissions", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &role_has_permissions_proto.CreateResponse{}, nil
}
