package login_attempts_controller

import (
	"context"
	"permission-service/generated_protos/login_attempts_proto"
	"permission-service/internal/app/ent"
	"permission-service/internal/app/ent/loginattempts"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *login_attempts_proto.GetRequest) (*login_attempts_proto.GetResponse, error) {
	login_attempts, err := c.Db.LoginAttempts.
		Query().
		Where(loginattempts.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.LoginAttemptsNotFound(int(in.Id))
	}

	return &login_attempts_proto.GetResponse{
		RequesterId: uint32(login_attempts.CreatedBy),
		UserId:      uint32(login_attempts.UserID),
		Successful:  bool(login_attempts.Successful),
	}, nil
}
