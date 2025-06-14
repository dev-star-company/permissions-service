package roles_controller

import (
	"context"
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
)

type Controller interface {
	roles_proto.ServiceServer

	Create(ctx context.Context, in *roles_proto.CreateRequest) (*roles_proto.CreateResponse, error)
	Get(ctx context.Context, in *roles_proto.GetRequest) (*roles_proto.GetResponse, error)
	Update(ctx context.Context, in *roles_proto.UpdateRequest) (*roles_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *roles_proto.DeleteRequest) (*roles_proto.DeleteResponse, error)
	List(ctx context.Context, in *roles_proto.ListRequest) (*roles_proto.ListResponse, error)
}

type controller struct {
	roles_proto.UnimplementedServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
