package first_login_controller

import (
	"context"
	"permission-service/generated_protos/first_login_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *first_login_proto.UpdateRequest) (*first_login_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.FirstLoginNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var first_login *ent.FirstLogin

	first_loginQ := tx.FirstLogin.UpdateOneID(int(in.Id))

	if in.UserId != nil && *in.UserId != 0 {
		first_loginQ.SetUserID(int(*in.UserId))
	}

	first_loginQ.SetUpdatedBy(int(in.RequesterId))

	first_login, err = first_loginQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.FirstLoginNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("first_login", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &first_login_proto.UpdateResponse{
		RequesterId: uint32(first_login.CreatedBy),
		UserId:      uint32(*first_login.UserID),
	}, nil
}
