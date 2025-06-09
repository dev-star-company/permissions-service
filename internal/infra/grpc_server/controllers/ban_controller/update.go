package ban_controller

import (
	"context"
	"permission-service/internal/app/ent"
	"permission-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permission-service/generated_protos/ban_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *ban_proto.UpdateRequest) (*ban_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.BanNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var ban *ent.Ban

	banQ := tx.Ban.UpdateOneID(int(in.Id))

	banQ.SetUpdatedBy(int(in.RequesterId))

	ban, err = banQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.BanNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("ban", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &ban_proto.UpdateResponse{
		RequesterId: uint32(ban.CreatedBy),
	}, nil
}
