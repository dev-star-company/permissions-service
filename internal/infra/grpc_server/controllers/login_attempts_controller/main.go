package login_attempts_controller

import (
	"context"
	"permission-service/generated_protos/login_attempts_proto"
	"permission-service/internal/app/ent"
)

type Controller interface {
	login_attempts_proto.LoginAttemptsServiceServer

	Create(ctx context.Context, in *login_attempts_proto.CreateRequest) (*login_attempts_proto.CreateResponse, error)
	Get(ctx context.Context, in *login_attempts_proto.GetRequest) (*login_attempts_proto.GetResponse, error)
	Update(ctx context.Context, in *login_attempts_proto.UpdateRequest) (*login_attempts_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *login_attempts_proto.DeleteRequest) (*login_attempts_proto.DeleteResponse, error)
	List(ctx context.Context, in *login_attempts_proto.ListRequest) (*login_attempts_proto.ListResponse, error)
}

type controller struct {
	login_attempts_proto.UnimplementedLoginAttemptsServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
