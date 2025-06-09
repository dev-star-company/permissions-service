package ban_controller

import (
	"context"
	"permission-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permission-service/generated_protos/ban_proto"
)

type Controller interface {
	ban_proto.BanServiceServer

	Create(ctx context.Context, in *ban_proto.CreateRequest) (*ban_proto.CreateResponse, error)
	Get(ctx context.Context, in *ban_proto.GetRequest) (*ban_proto.GetResponse, error)
	Update(ctx context.Context, in *ban_proto.UpdateRequest) (*ban_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *ban_proto.DeleteRequest) (*ban_proto.DeleteResponse, error)
	List(ctx context.Context, in *ban_proto.ListRequest) (*ban_proto.ListResponse, error)
}

type controller struct {
	ban_proto.UnimplementedBanServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
