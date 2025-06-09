package role_has_permissions_controller

import (
	"context"
	"permission-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permission-service/generated_protos/role_has_permissions_proto"
)

type Controller interface {
	role_has_permissions_proto.RoleHasPermissionsServiceServer

	Create(ctx context.Context, in *role_has_permissions_proto.CreateRequest) (*role_has_permissions_proto.CreateResponse, error)
	Get(ctx context.Context, in *role_has_permissions_proto.GetRequest) (*role_has_permissions_proto.GetResponse, error)
	Update(ctx context.Context, in *role_has_permissions_proto.UpdateRequest) (*role_has_permissions_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *role_has_permissions_proto.DeleteRequest) (*role_has_permissions_proto.DeleteResponse, error)
	List(ctx context.Context, in *role_has_permissions_proto.ListRequest) (*role_has_permissions_proto.ListResponse, error)
}

type controller struct {
	role_has_permissions_proto.UnimplementedRoleHasPermissionsServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
