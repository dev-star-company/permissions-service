package first_login_controller

import (
	"context"
	"permission-service/generated_protos/first_login_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *first_login_proto.CreateRequest) (*first_login_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.FirstLogin.Create().
		SetUserID(int(in.UserId)).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &first_login_proto.CreateResponse{
		RequesterId: uint32(create.CreatedBy),
		UserId:      uint32(*create.UserID),
	}, nil
}
