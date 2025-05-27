package roles_controller

import (
	"context"
	"permission-service/generated_protos/roles_proto"
	"permission-service/internal/adapters/grpc_controllers"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *roles_proto.UpdateRequest) (*roles_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	roleQ := tx.Role.
		UpdateOneID(int(in.Id)).
		SetUpdatedBy(int(in.RequesterId))

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
		Role: grpc_controllers.RoleToProto(role),
	}, nil
}
