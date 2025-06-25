package role_has_permissions_controller

import (
	"context"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *role_has_permissions_proto.CreateRequest) (*role_has_permissions_proto.CreateResponse, error) {

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

	_, err = c.Db.RoleHasPermissions.Create().
		SetCreatedBy(requesterId).
		SetUpdatedBy(requesterId).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("role_has_permissions", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &role_has_permissions_proto.CreateResponse{
		RequesterUuid: in.RequesterUuid,
	}, nil
}
