package service_controller

import (
	"context"
	"permission-service/generated_protos/service_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *service_proto.UpdateRequest) (*service_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ServiceNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var service *ent.Services

	serviceQ := tx.Services.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		serviceQ.SetName(string(*in.Name))
	}

	serviceQ.SetUpdatedBy(int(in.RequesterId))

	service, err = serviceQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ServiceNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("service", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &service_proto.UpdateResponse{
		RequesterId: uint32(service.CreatedBy),
		Name:        string(service.Name),
	}, nil
}
