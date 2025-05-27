package users_controller

import (
	"context"
	"permission-service/generated_protos/users_proto"
	"permission-service/internal/app/ent"
)

type Controller interface {
	users_proto.UsersServiceServer

	Create(ctx context.Context, in *users_proto.CreateRequest) (*users_proto.CreateResponse, error)
	Get(ctx context.Context, in *users_proto.GetRequest) (*users_proto.GetResponse, error)
	Update(ctx context.Context, in *users_proto.UpdateRequest) (*users_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *users_proto.DeleteRequest) (*users_proto.DeleteResponse, error)
	List(ctx context.Context, in *users_proto.ListRequest) (*users_proto.ListResponse, error)

	VerifyPassword(ctx context.Context, in *users_proto.VerifyPasswordRequest) (*users_proto.VerifyPasswordResponse, error)
}

type controller struct {
	users_proto.UnimplementedUsersServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
