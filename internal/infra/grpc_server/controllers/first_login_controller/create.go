package first_login_controller

import (
	"context"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *first_login_proto.CreateRequest) (*first_login_proto.CreateResponse, error) {

	if in.RequesterUuid == "" {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	requesterId, err := controllers.GetRequesterId(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	create, err := c.Db.FirstLogin.Create().
		SetUserID(int(in.UserId)).
		SetCreatedBy(requesterId).
		SetUpdatedBy(requesterId).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("first_login", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &first_login_proto.CreateResponse{
		RequesterUuid: in.RequesterUuid,
		UserId:        uint32(*create.UserID),
	}, nil
}
