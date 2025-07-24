package permission_controller

import (
	"context"

	"permissions-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *permission_proto.DeleteRequest) (*permission_proto.DeleteResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.Permission.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("permission", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &permission_proto.DeleteResponse{}, nil
}
