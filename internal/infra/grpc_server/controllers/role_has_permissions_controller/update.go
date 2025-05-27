package role_has_permissions_controller

import (
	"context"
	"permission-service/generated_protos/role_has_permissions_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *role_has_permissions_proto.UpdateRequest) (*role_has_permissions_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.RoleHasPermissionNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var role_has_permissions *ent.RoleHasPermissions

	role_has_permissionsQ := tx.RoleHasPermissions.UpdateOneID(int(in.Id))

	role_has_permissionsQ.SetUpdatedBy(int(in.RequesterId))

	role_has_permissions, err = role_has_permissionsQ.Save(ctx)
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
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &role_has_permissions_proto.UpdateResponse{
		RequesterId: uint32(role_has_permissions.CreatedBy),
	}, nil
}
