package ban_controller

import (
	"context"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/ban_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *ban_proto.DeleteRequest) (*ban_proto.DeleteResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.Ban.UpdateOneID(int(in.Id)).Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("ban", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &ban_proto.DeleteResponse{}, nil
}
