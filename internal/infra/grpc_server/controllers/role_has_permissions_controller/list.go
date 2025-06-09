package role_has_permissions_controller

import (
	"context"
	"errors"
	"permissions-service/internal/adapters/grpc_controllers"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/rolehaspermissions"
	"permissions-service/internal/app/ent/schema"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *role_has_permissions_proto.ListRequest) (*role_has_permissions_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.RoleHasPermissions.Query()

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying role_has_permissions", err)
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
				query = query.Order(ent.Asc(rolehaspermissions.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(rolehaspermissions.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	role_has_permissions, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying role_has_permissions", err)
	}

	responseRoleHasPermissions := make([]*role_has_permissions_proto.RoleHasPermissions, len(role_has_permissions))
	for i, acc := range role_has_permissions {
		responseRoleHasPermissions[i] = grpc_controllers.RoleHasPermissionsToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &role_has_permissions_proto.ListResponse{
		Rows:  responseRoleHasPermissions,
		Count: uint32(count),
	}, nil
}
