package roles_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
)

func (c *controller) Delete(ctx context.Context, in *roles_proto.DeleteRequest) (*roles_proto.DeleteResponse, error) {

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
