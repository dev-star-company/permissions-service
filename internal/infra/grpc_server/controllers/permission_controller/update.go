package permission_controller

import (
	"context"
	"permission-service/generated_protos/permission_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *permission_proto.UpdateRequest) (*permission_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.PermissionNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var permission *ent.Permission

	permissionQ := tx.Permission.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		permissionQ.SetName(string(*in.Name))
	}

	permissionQ.SetUpdatedBy(int(in.RequesterId))

	permission, err = permissionQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.PermissionNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("permission", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &permission_proto.UpdateResponse{
		Name:         permission.Name,
		Description:  permission.Description,
		InternalName: permission.InternalName,
		IsActive:     permission.IsActive,
		ServiceId:    uint32(permission.ServiceID),
	}, nil
}
