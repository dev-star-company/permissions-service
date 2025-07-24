package first_login_controller

import (
	"context"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *first_login_proto.CreateRequest) (*first_login_proto.CreateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	defer tx.Rollback()

	user, err := c.Db.User.Get(ctx, int(in.UserId))
	if err != nil {
		return nil, err
	}

	create, err := tx.FirstLogin.Create().
		SetUserID(user.ID).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("first_login", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &first_login_proto.CreateResponse{
		UserId: int32(create.Edges.User.ID),
	}, nil
}
