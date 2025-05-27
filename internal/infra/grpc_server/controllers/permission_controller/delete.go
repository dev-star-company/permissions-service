package permission_controller

import (
	"context"
	"permission-service/generated_protos/permission_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *permission_proto.DeleteRequest) (*permission_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.PermissionNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.Permission.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("permission", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &permission_proto.DeleteResponse{}, nil
}
