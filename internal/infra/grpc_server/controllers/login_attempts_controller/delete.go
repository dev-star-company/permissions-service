package login_attempts_controller

import (
	"context"
	"permission-service/generated_protos/login_attempts_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *login_attempts_proto.DeleteRequest) (*login_attempts_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.LoginAttemptsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.LoginAttempts.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("login_attempts", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &login_attempts_proto.DeleteResponse{}, nil
}
