package first_login_controller

import (
	"context"
	"permission-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permission-service/generated_protos/first_login_proto"
)

type Controller interface {
	first_login_proto.FirstLoginServiceServer

	Create(ctx context.Context, in *first_login_proto.CreateRequest) (*first_login_proto.CreateResponse, error)
	Get(ctx context.Context, in *first_login_proto.GetRequest) (*first_login_proto.GetResponse, error)
	Update(ctx context.Context, in *first_login_proto.UpdateRequest) (*first_login_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *first_login_proto.DeleteRequest) (*first_login_proto.DeleteResponse, error)
	List(ctx context.Context, in *first_login_proto.ListRequest) (*first_login_proto.ListResponse, error)
}

type controller struct {
	first_login_proto.UnimplementedFirstLoginServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
