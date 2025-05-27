package service_controller

import (
	"context"
	"permission-service/generated_protos/service_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *service_proto.DeleteRequest) (*service_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ServiceNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.Services.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("service", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &service_proto.DeleteResponse{}, nil
}
