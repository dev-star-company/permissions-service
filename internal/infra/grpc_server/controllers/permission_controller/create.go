package permission_controller

import (
	"context"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *permission_proto.CreateRequest) (*permission_proto.CreateResponse, error) {

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

	create, err := c.Db.Permission.Create().
		SetName(in.Name).
		SetCreatedBy(requester.ID).
		SetUpdatedBy(requester.ID).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("permission", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &permission_proto.CreateResponse{
		Name:         create.Name,
		Description:  create.Description,
		InternalName: create.InternalName,
		IsActive:     create.IsActive,
		ServiceId:    uint32(create.ServiceID),
	}, nil
}
