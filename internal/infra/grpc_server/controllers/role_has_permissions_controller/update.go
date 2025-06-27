package role_has_permissions_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *role_has_permissions_proto.UpdateRequest) (*role_has_permissions_proto.UpdateResponse, error) {
	if in.RequesterUuid == "" {
		return nil, errs.RoleHasPermissionNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	role_has_permissionsQ := tx.RoleHasPermissions.UpdateOneID(int(in.Id))

	role_has_permissionsQ.SetUpdatedBy(requester.ID)

	_, err = role_has_permissionsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.RoleHasPermissionNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("role_has_permissions", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &role_has_permissions_proto.UpdateResponse{
		RequesterUuid: in.RequesterUuid,
	}, nil
}
