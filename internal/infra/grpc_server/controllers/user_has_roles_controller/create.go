package user_has_roles_controller

import (
	"context"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/user_has_roles_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *user_has_roles_proto.CreateRequest) (*user_has_roles_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	_, err = c.Db.UserHasRoles.Create().
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("user_has_roles", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &user_has_roles_proto.CreateResponse{}, nil
}
