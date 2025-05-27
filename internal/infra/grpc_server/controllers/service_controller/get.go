package service_controller

import (
	"context"
	"permission-service/generated_protos/service_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/app/ent/services"
	"permission-service/internal/pkg/errs"
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
