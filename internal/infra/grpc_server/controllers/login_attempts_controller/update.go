package login_attempts_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/login_attempts_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *login_attempts_proto.UpdateRequest) (*login_attempts_proto.UpdateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	login_attemptsQ := tx.LoginAttempts.UpdateOneID(int(in.Id))

	if in.UserUuid != nil && *in.UserUuid != "" {
		user, err := controllers.GetUserFromUuid(tx, ctx, *in.UserUuid)
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}
		login_attemptsQ.SetUserID(user.ID)
	}

	if in.Successful != nil {
		login_attemptsQ.SetSuccessful(bool(*in.Successful))
	}

	login_attemptsQ.SetUpdatedBy(requester.ID)

	login_attempts, err := login_attemptsQ.Save(ctx)
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
		RequesterUuid: in.RequesterUuid,
		UserUuid:      login_attempts.Edges.User.UUID.String(),
		Successful:    bool(login_attempts.Successful),
	}, nil
}
