package permission_controller

import (
	"context"
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"
)

type Controller interface {
	permission_proto.PermissionServiceServer

	Create(ctx context.Context, in *permission_proto.CreateRequest) (*permission_proto.CreateResponse, error)
	Get(ctx context.Context, in *permission_proto.GetRequest) (*permission_proto.GetResponse, error)
	Update(ctx context.Context, in *permission_proto.UpdateRequest) (*permission_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *permission_proto.DeleteRequest) (*permission_proto.DeleteResponse, error)
	List(ctx context.Context, in *permission_proto.ListRequest) (*permission_proto.ListResponse, error)
}

type controller struct {
	permission_proto.UnimplementedPermissionServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
