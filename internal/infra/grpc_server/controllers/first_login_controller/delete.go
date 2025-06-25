package first_login_controller

import (
	"context"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *first_login_proto.DeleteRequest) (*first_login_proto.DeleteResponse, error) {
	if in.RequesterUuid == "" {
		return nil, errs.FirstLoginNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	requesterId, err := controllers.GetRequesterId(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	err = tx.FirstLogin.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(requesterId).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("first_login", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &first_login_proto.DeleteResponse{}, nil
}
