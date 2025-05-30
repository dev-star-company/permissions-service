package ban_controller

import (
	"context"
	"errors"
	"permission-service/generated_protos/ban_proto"
	"permission-service/internal/adapters/grpc_controllers"
	"permission-service/internal/app/ent"
	"permission-service/internal/app/ent/ban"
	"permission-service/internal/app/ent/schema"
	"permission-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *ban_proto.ListRequest) (*ban_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Ban.Query()

	if in.ExpiresAt != nil && *in.ExpiresAt != "" {
		parsedTime, err := time.Parse(time.RFC3339, *in.ExpiresAt)
		if err != nil {
			return nil, errs.ListingError("querying ban", err)
		}
		query = query.Where(ban.ExpiresAtEQ(parsedTime))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying ban", err)
	}

	if in.Limit != nil && *in.Limit > 0 {
		query = query.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		query = query.Offset(int(*in.Offset))
	}

	if in.Orderby != nil {
		if in.Orderby.Id != nil {
			switch *in.Orderby.Id {
			case "ASC":
				query = query.Order(ent.Asc(ban.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(ban.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	ban, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying ban", err)
	}

	responseBan := make([]*ban_proto.Ban, len(ban))
	for i, acc := range ban {
		responseBan[i] = grpc_controllers.BanToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &ban_proto.ListResponse{
		Rows:  responseBan,
		Count: uint32(count),
	}, nil
}
