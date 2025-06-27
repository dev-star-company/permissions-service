package login_attempts_controller

import (
	"context"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/login_attempts_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *login_attempts_proto.CreateRequest) (*login_attempts_proto.CreateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	user, err := controllers.GetUserFromUuid(tx, ctx, in.UserUuid)
	if err != nil {
		return nil, err
	}

	create, err := c.Db.LoginAttempts.Create().
		SetUserID(user.ID).
		SetSuccessful(in.Successful).
		SetCreatedBy(requester.ID).
		SetUpdatedBy(requester.ID).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("login_attempts", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &login_attempts_proto.CreateResponse{
		RequesterUuid: in.RequesterUuid,
		UserUuid:      user.UUID.String(),
		Successful:    bool(create.Successful),
	}, nil
}
