package service_controller

import (
	"context"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/service_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *service_proto.CreateRequest) (*service_proto.CreateResponse, error) {

	if in.RequesterUuid == "" {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	requester, err := controllers.GetUserIdFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	create, err := c.Db.Services.Create().
		SetName(in.Name).
		SetCreatedBy(requester.ID).
		SetUpdatedBy(requester.ID).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("service", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &service_proto.CreateResponse{
		RequesterUuid: in.RequesterUuid,
		Name:          string(create.Name),
	}, nil
}
