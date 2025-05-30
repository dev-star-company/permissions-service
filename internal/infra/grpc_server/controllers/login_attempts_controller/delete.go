package login_attempts_controller

import (
	"context"
	"permission-service/generated_protos/login_attempts_proto"
	"permission-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *login_attempts_proto.DeleteRequest) (*login_attempts_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.LoginAttemptsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.LoginAttempts.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("login_attempts", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &login_attempts_proto.DeleteResponse{}, nil
}
