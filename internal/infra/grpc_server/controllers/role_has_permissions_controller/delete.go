package role_has_permissions_controller

import (
	"context"

	"permissions-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *role_has_permissions_proto.DeleteRequest) (*role_has_permissions_proto.DeleteResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.RoleHasPermissions.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("role_has_permissions", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &role_has_permissions_proto.DeleteResponse{}, nil
}
