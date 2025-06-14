package permission_controller

import (
	"context"
	"errors"
	"permissions-service/internal/adapters/grpc_controllers"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/permission"
	"permissions-service/internal/app/ent/schema"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *permission_proto.ListRequest) (*permission_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Permission.Query()

	if in.Name != nil {
		query = query.Where(permission.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying permission", err)
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
				query = query.Order(ent.Asc(permission.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(permission.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.OrderBy.Id))
			}
		}
	}

	permission, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying permission", err)
	}

	responsePermission := make([]*permission_proto.Permission, len(permission))
	for i, acc := range permission {
		responsePermission[i] = grpc_controllers.PermissionToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &permission_proto.ListResponse{
		Rows:  responsePermission,
		Count: uint32(count),
	}, nil
}
