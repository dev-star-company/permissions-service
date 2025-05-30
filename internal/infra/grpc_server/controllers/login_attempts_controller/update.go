package login_attempts_controller

import (
	"context"
	"permission-service/generated_protos/login_attempts_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *login_attempts_proto.UpdateRequest) (*login_attempts_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.LoginAttemptsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var login_attempts *ent.LoginAttempts

	login_attemptsQ := tx.LoginAttempts.UpdateOneID(int(in.Id))

	if in.UserId != nil && *in.UserId != 0 {
		login_attemptsQ.SetUserID(int(*in.UserId))
	}

	if in.Successful != nil && *in.Successful {
		login_attemptsQ.SetSuccessful(bool(*in.Successful))
	}

	login_attemptsQ.SetUpdatedBy(int(in.RequesterId))

	login_attempts, err = login_attemptsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.LoginAttemptsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("login_attempts", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &login_attempts_proto.UpdateResponse{
		RequesterId: uint32(login_attempts.CreatedBy),
		UserId:      uint32(login_attempts.UserID),
		Successful:  bool(login_attempts.Successful),
	}, nil
}
