package first_login_controller

import (
	"context"
	"permission-service/internal/app/ent"
	"permission-service/internal/app/ent/firstlogin"

	"github.com/dev-star-company/protos-go/permission-service/generated_protos/first_login_proto"

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
		RequesterId: uint32(first_login.CreatedBy),
		UserId:      uint32(*first_login.UserID),
	}, nil
}
