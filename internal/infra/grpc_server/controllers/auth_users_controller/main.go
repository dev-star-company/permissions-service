package auth_users_controller

import (
	"context"
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
)

type Controller interface {
	auth_users_proto.AuthUsersServiceServer

	Create(ctx context.Context, in *auth_users_proto.CreateRequest) (*auth_users_proto.CreateResponse, error)
	Get(ctx context.Context, in *auth_users_proto.GetRequest) (*auth_users_proto.GetResponse, error)
	GetUserRoles(ctx context.Context, in *auth_users_proto.GetUserByRolesRequest) (*auth_users_proto.GetUserByRolesResponse, error)
	Update(ctx context.Context, in *auth_users_proto.UpdateRequest) (*auth_users_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *auth_users_proto.DeleteRequest) (*auth_users_proto.DeleteResponse, error)
	List(ctx context.Context, in *auth_users_proto.ListRequest) (*auth_users_proto.ListResponse, error)

	VerifyPassword(ctx context.Context, in *auth_users_proto.VerifyPasswordRequest) (*auth_users_proto.VerifyPasswordResponse, error)
}

type controller struct {
	auth_users_proto.UnimplementedAuthUsersServiceServer

	Db *ent.Client
	k  *connection.Connectioner
}

func New(Db *ent.Client, k *connection.Connectioner) Controller {
	return &controller{
		Db: Db,
		k:  k,
	}
}
