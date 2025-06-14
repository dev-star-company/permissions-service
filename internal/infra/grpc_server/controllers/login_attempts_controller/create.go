package login_attempts_controller

import (
	"context"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/login_attempts_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *login_attempts_proto.CreateRequest) (*login_attempts_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	create, err := c.Db.LoginAttempts.Create().
		SetUserID(int(in.UserId)).
		SetSuccessful(in.Successful).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &login_attempts_proto.CreateResponse{
		RequesterId: uint32(create.CreatedBy),
		UserId:      uint32(create.UserID),
		Successful:  bool(create.Successful),
	}, nil
}
