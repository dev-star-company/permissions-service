package role_has_permissions_controller

import (
	"context"
	"permission-service/generated_protos/role_has_permissions_proto"
	"permission-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *role_has_permissions_proto.DeleteRequest) (*role_has_permissions_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.RoleHasPermissionNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.RoleHasPermissions.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("role_has_permissions", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &role_has_permissions_proto.DeleteResponse{}, nil
}
