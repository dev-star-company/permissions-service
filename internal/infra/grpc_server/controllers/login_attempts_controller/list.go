package login_attempts_controller

import (
	"context"
	"errors"
	"permission-service/generated_protos/login_attempts_proto"
	"permission-service/internal/adapters/grpc_controllers"
	"permission-service/internal/app/ent"
	"permission-service/internal/app/ent/loginattempts"
	"permission-service/internal/app/ent/schema"
	"permission-service/internal/pkg/errs"
	"permission-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *login_attempts_proto.ListRequest) (*login_attempts_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.LoginAttempts.Query()

	if in.UserId != nil {
		query = query.Where(loginattempts.UserID(int(*in.UserId)))
	}

	if in.Successful != nil {
		query = query.Where(loginattempts.Successful(bool(*in.Successful)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying login_attempts", err)
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
				query = query.Order(ent.Asc(loginattempts.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(loginattempts.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	login_attempts, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying login_attempts", err)
	}

	responseLoginAttempts := make([]*login_attempts_proto.LoginAttempts, len(login_attempts))
	for i, acc := range login_attempts {
		responseLoginAttempts[i] = grpc_controllers.LoginAttemptsToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &login_attempts_proto.ListResponse{
		Rows:  responseLoginAttempts,
		Count: uint32(count),
	}, nil
}
