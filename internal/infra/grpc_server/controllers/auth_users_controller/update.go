package auth_users_controller

import (
	"context"
	"permissions-service/internal/adapters/grpc_convertions"
	userSchema "permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *auth_users_proto.UpdateRequest) (*auth_users_proto.UpdateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.Id == nil || *in.Id == 0 {
		return nil, errs.RequesterIDRequired()
	}

	target, err := c.Db.User.Get(ctx, int(*in.Id))
	if err != nil {
		return nil, err
	}

	userQ := tx.User.UpdateOneID(int(target.ID))

	if in.Name != nil {
		userQ.SetName(*in.Name)
	}
	if in.Surname != nil {
		userQ.SetSurname(*in.Surname)
	}

	user, err := userQ.Save(ctx)
	if err != nil {
		return nil, utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, err)
	}

	userWithRelations, err := c.Db.User.Query().
		Where(userSchema.IDEQ(user.ID)).
		WithPhones().
		WithEmails().
		WithRoles().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &auth_users_proto.UpdateResponse{
		User: grpc_convertions.UserToProto(userWithRelations),
	}, nil
}
