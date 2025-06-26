package first_login_controller

import (
	"context"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *first_login_proto.CreateRequest) (*first_login_proto.CreateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	defer tx.Rollback()

	requester, err := controllers.GetUserIdFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	user, err := controllers.GetUserIdFromUuid(tx, ctx, in.UserUuid)
	if err != nil {
		return nil, err
	}

	create, err := tx.FirstLogin.Create().
		SetUserID(user.ID).
		SetCreatedBy(requester.ID).
		SetUpdatedBy(requester.ID).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("first_login", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &first_login_proto.CreateResponse{
		RequesterUuid: in.RequesterUuid,
		UserUuid:      create.Edges.User.UUID.String(),
	}, nil
}
