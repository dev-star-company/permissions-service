package service_controller

import (
	"context"
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/service_proto"
)

type Controller interface {
	service_proto.ServiceServer

	Create(ctx context.Context, in *service_proto.CreateRequest) (*service_proto.CreateResponse, error)
	Get(ctx context.Context, in *service_proto.GetRequest) (*service_proto.GetResponse, error)
	Update(ctx context.Context, in *service_proto.UpdateRequest) (*service_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *service_proto.DeleteRequest) (*service_proto.DeleteResponse, error)
	List(ctx context.Context, in *service_proto.ListRequest) (*service_proto.ListResponse, error)
}

type controller struct {
	service_proto.UnimplementedServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
