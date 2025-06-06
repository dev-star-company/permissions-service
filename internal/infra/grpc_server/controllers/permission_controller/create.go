package permission_controller

import (
	"context"
	"permission-service/generated_protos/permission_proto"
	"permission-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *permission_proto.CreateRequest) (*permission_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	create, err := c.Db.Permission.Create().
		SetName(in.Name).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &permission_proto.CreateResponse{
		Name:         create.Name,
		Description:  create.Description,
		InternalName: create.InternalName,
		IsActive:     create.IsActive,
		ServiceId:    uint32(create.ServiceID),
	}, nil
}
