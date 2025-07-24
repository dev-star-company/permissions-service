package ban_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/pkg/utils"
	"strconv"
	"time"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/ban_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *ban_proto.CreateRequest) (*ban_proto.CreateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}
	defer tx.Rollback()

	expiresAtInt, err := strconv.ParseInt(in.ExpiresAt, 10, 64)
	if err != nil {
		return nil, errs.BadRequest(fmt.Errorf("invalid expires_at: %w", err))
	}
	expiresAt := time.Unix(expiresAtInt, 0)

	_, err = c.Db.Ban.Create().
		SetExpiresAt(expiresAt).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("ban", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &ban_proto.CreateResponse{}, nil
}
