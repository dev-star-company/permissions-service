package service_controller

import (
	"context"
	"permission-service/generated_protos/service_proto"
	"permission-service/internal/app/ent"
)

type Controller interface {
	service_proto.ServiceServiceServer

	Create(ctx context.Context, in *service_proto.CreateRequest) (*service_proto.CreateResponse, error)
	Get(ctx context.Context, in *service_proto.GetRequest) (*service_proto.GetResponse, error)
	Update(ctx context.Context, in *service_proto.UpdateRequest) (*service_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *service_proto.DeleteRequest) (*service_proto.DeleteResponse, error)
	List(ctx context.Context, in *service_proto.ListRequest) (*service_proto.ListResponse, error)
}

type controller struct {
	service_proto.UnimplementedServiceServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
