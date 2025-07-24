package first_login_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/firstlogin"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *first_login_proto.GetRequest) (*first_login_proto.GetResponse, error) {
	first_login, err := c.Db.FirstLogin.
		Query().
		Where(firstlogin.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.FirstLoginNotFound(int(in.Id))
	}

	return &first_login_proto.GetResponse{
		UserId: int32(*first_login.UserID),
	}, nil
}
