package roles_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/infra/grpc_server/controllers"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
)

func (c *controller) Create(ctx context.Context, in *roles_proto.CreateRequest) (*roles_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	// Create a new role in the database
	role, err := tx.Role.Create().
		SetName(in.Name).
		SetDescription(in.Description).
		SetCreatedBy(requester.ID).
		SetUpdatedBy(requester.ID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return &roles_proto.CreateResponse{
		Role: grpc_convertions.RoleToProto(role),
	}, nil
}
