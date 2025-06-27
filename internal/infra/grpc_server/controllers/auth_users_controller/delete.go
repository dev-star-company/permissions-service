package auth_users_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *auth_users_proto.DeleteRequest) (*auth_users_proto.DeleteResponse, error) {
	if in.RequesterUuid == "" {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	defer tx.Rollback()

	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	// Delete the user from the database
	err = tx.User.DeleteOneID(int(in.Id)).Exec(ctx)
	if err != nil {
		return nil, utils.Rollback(tx, fmt.Errorf("deleting user: %w", err))
	}

	err = tx.User.UpdateOneID(int(in.Id)).SetDeletedBy(requester.ID).Exec(ctx)
	if err != nil {
		return nil, utils.Rollback(tx, fmt.Errorf("deleting user: %w", err))
	}
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	// Create and return the response
	return &auth_users_proto.DeleteResponse{}, nil
}
