package ban_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/ban"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/ban_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *ban_proto.GetRequest) (*ban_proto.GetResponse, error) {
	ban, err := c.Db.Ban.
		Query().
		Where(ban.ID(int(in.Id))).
		WithUser().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.BanNotFound(int(in.Id))
	}

	return &ban_proto.GetResponse{
		ExpiresAt: ban.ExpiresAt.String(),
	}, nil
}
