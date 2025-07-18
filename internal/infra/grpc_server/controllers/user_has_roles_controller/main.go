package user_has_roles_controller

import (
	"context"
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/user_has_roles_proto"
)

type Controller interface {
	user_has_roles_proto.UserHasRolesServiceServer

	Create(ctx context.Context, in *user_has_roles_proto.CreateRequest) (*user_has_roles_proto.CreateResponse, error)
	Get(ctx context.Context, in *user_has_roles_proto.GetRequest) (*user_has_roles_proto.GetResponse, error)
	Update(ctx context.Context, in *user_has_roles_proto.UpdateRequest) (*user_has_roles_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *user_has_roles_proto.DeleteRequest) (*user_has_roles_proto.DeleteResponse, error)
	List(ctx context.Context, in *user_has_roles_proto.ListRequest) (*user_has_roles_proto.ListResponse, error)
}

type controller struct {
	user_has_roles_proto.UserHasRolesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
