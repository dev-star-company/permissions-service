package user_has_roles_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/user_has_roles_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *user_has_roles_proto.UpdateRequest) (*user_has_roles_proto.UpdateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	user_has_rolesQ := tx.UserHasRoles.UpdateOneID(int(in.Id))

	_, err = user_has_rolesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.UserHasRolesNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("user_has_roles", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &user_has_roles_proto.UpdateResponse{}, nil
}
