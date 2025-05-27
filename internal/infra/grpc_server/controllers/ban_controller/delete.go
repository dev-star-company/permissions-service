package ban_controller

import (
	"context"
	"permission-service/generated_protos/ban_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *ban_proto.DeleteRequest) (*ban_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.BanNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.Ban.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("ban", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &ban_proto.DeleteResponse{}, nil
}
