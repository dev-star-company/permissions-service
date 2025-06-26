package roles_controller

import (
	"context"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *roles_proto.UpdateRequest) (*roles_proto.UpdateResponse, error) {
	if in.RequesterUuid == "" {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	requester, err := controllers.GetUserIdFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	roleQ := tx.Role.
		UpdateOneID(int(in.Id)).
		SetUpdatedBy(requester.ID)

	if in.Name != nil {
		roleQ.SetName(*in.Name)
	}

	if in.Description != nil {
		roleQ.SetDescription(*in.Description)
	}

	role, err := roleQ.Save(ctx)
	if err != nil {
		return nil, utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, err)
	}

	return &roles_proto.UpdateResponse{
		Role: grpc_convertions.RoleToProto(role),
	}, nil
}
