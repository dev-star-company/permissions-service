package ban_controller

import (
	"context"
	"permission-service/generated_protos/ban_proto"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
	"time"
)

func (c *controller) Create(ctx context.Context, in *ban_proto.CreateRequest) (*ban_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	expiresAt, err := time.Parse(time.RFC3339, in.ExpiresAt)
	if err != nil {
		return nil, errs.CreateError("expires_at", err)
	}

	create, err := c.Db.Ban.Create().
		SetUserID(int(in.UserId)).
		SetExpiresAt(expiresAt).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &ban_proto.CreateResponse{
		RequesterId: uint32(create.CreatedBy),
		UserId:      uint32(create.UserID),
		ExpiresAt: func() string {
			if create.ExpiresAt != nil {
				return create.ExpiresAt.Format(time.RFC3339)
			}
			return ""
		}(),
	}, nil
}
