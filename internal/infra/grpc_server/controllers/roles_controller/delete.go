package roles_controller

import (
	"context"
	"fmt"
	"permission-service/generated_protos/roles_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) Delete(ctx context.Context, in *roles_proto.DeleteRequest) (*roles_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	// Delete the role from the database
	err = tx.Role.DeleteOneID(int(in.Id)).Exec(ctx)
	if err != nil {
		return nil, utils.Rollback(tx, fmt.Errorf("deleting role: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, fmt.Errorf("committing transaction: %w", err))
	}

	return &roles_proto.DeleteResponse{}, nil
}
