package users_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *auth_users_proto.DeleteRequest) (*auth_users_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	// Delete the user from the database
	err = tx.User.DeleteOneID(int(in.Id)).Exec(ctx)
	if err != nil {
		return nil, utils.Rollback(tx, fmt.Errorf("deleting user: %w", err))
	}

	err = tx.User.UpdateOneID(int(in.Id)).SetDeletedBy(int(in.RequesterId)).Exec(ctx)
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
