package service_controller

import (
	"context"
	"errors"
	"permissions-service/internal/adapters/grpc_controllers"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/schema"
	"permissions-service/internal/app/ent/services"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/service_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *service_proto.ListRequest) (*service_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Services.Query()

	if in.Name != nil {
		query = query.Where(services.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying service", err)
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
				query = query.Order(ent.Asc(services.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(services.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.OrderBy.Id))
			}
		}
	}

	service, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying service", err)
	}

	responseService := make([]*service_proto.ServiceDto, len(service))
	for i, acc := range service {
		responseService[i] = grpc_controllers.ServiceToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &service_proto.ListResponse{
		Rows:  responseService,
		Count: uint32(count),
	}, nil
}
