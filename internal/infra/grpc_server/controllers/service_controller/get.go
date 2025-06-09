package service_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/services"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/service_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *service_proto.GetRequest) (*service_proto.GetResponse, error) {
	service, err := c.Db.Services.
		Query().
		Where(services.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ServiceNotFound(int(in.Id))
	}

	return &service_proto.GetResponse{
		RequesterId: uint32(service.CreatedBy),
		Name:        service.Name,
	}, nil
}
