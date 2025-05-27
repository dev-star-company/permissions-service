package ban_controller

import (
	"context"
	"permission-service/generated_protos/ban_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/app/ent/ban"
	"permission-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *ban_proto.GetRequest) (*ban_proto.GetResponse, error) {
	ban, err := c.Db.Ban.
		Query().
		Where(ban.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.BanNotFound(int(in.Id))
	}

	return &ban_proto.GetResponse{
		RequesterId: uint32(ban.CreatedBy),
	}, nil
}
