package permission_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *permission_proto.UpdateRequest) (*permission_proto.UpdateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	permissionQ := tx.Permission.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		permissionQ.SetName(string(*in.Name))
	}

	permission, err := permissionQ.Save(ctx)
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
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &permission_proto.UpdateResponse{
		Name:         permission.Name,
		Description:  permission.Description,
		InternalName: permission.InternalName,
		IsActive:     permission.IsActive,
		ServiceId:    uint32(permission.ServiceID),
	}, nil
}
