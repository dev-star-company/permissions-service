package password_controller

import (
	"context"
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/password_proto"
)

type Controller interface {
	password_proto.PasswordServiceServer

	Verify(ctx context.Context, in *password_proto.VerifyRequest) (*password_proto.VerifyResponse, error)
}

type controller struct {
	password_proto.UnimplementedPasswordServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
