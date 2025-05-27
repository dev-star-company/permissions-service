package login_attempts_controller

import (
	"context"
	"permission-service/generated_protos/login_attempts_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *login_attempts_proto.CreateRequest) (*login_attempts_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
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
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &login_attempts_proto.CreateResponse{
		RequesterId: uint32(create.CreatedBy),
		UserId:      uint32(create.UserID),
		Successful:  bool(create.Successful),
	}, nil
}
