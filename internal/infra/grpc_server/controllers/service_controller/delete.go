package service_controller

import (
	"context"
	"permission-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/permission-service/generated_protos/service_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *service_proto.DeleteRequest) (*service_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ServiceNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.Services.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("service", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &service_proto.DeleteResponse{}, nil
}
