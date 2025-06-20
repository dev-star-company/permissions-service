package first_login_controller

import (
	"context"
	"errors"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/firstlogin"
	"permissions-service/internal/app/ent/schema"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *first_login_proto.ListRequest) (*first_login_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.FirstLogin.Query()

	if in.UserId != nil {
		query = query.Where(firstlogin.UserID(int(*in.UserId)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying first_login", err)
	}

	if in.Limit != nil && *in.Limit > 0 {
		query = query.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		query = query.Offset(int(*in.Offset))
	}

	if in.OrderBy != nil {
		if in.OrderBy.Id != nil {
			switch *in.OrderBy.Id {
			case "ASC":
				query = query.Order(ent.Asc(firstlogin.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(firstlogin.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.OrderBy.Id))
			}
		}
	}

	first_login, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying first_login", err)
	}

	responseFirstLogin := make([]*first_login_proto.FirstLogin, len(first_login))
	for i, acc := range first_login {
		responseFirstLogin[i] = grpc_convertions.FirstLoginToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &first_login_proto.ListResponse{
		Rows:  responseFirstLogin,
		Count: uint32(count),
	}, nil
}
