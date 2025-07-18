package user_has_roles_controller

import (
	"context"
	"errors"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/schema"
	"permissions-service/internal/app/ent/userhasroles"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/user_has_roles_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *user_has_roles_proto.ListRequest) (*user_has_roles_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.UserHasRoles.Query()

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying user_has_roles", err)
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
				query = query.Order(ent.Asc(userhasroles.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(userhasroles.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.OrderBy.Id))
			}
		}
	}

	user_has_roles, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying user_has_roles", err)
	}

	responseUserHasRoles := make([]*user_has_roles_proto.UserHasRoles, len(user_has_roles))
	for i, acc := range user_has_roles {
		responseUserHasRoles[i] = grpc_convertions.UserHasRolesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &user_has_roles_proto.ListResponse{
		Rows:  responseUserHasRoles,
		Count: uint32(count),
	}, nil
}
