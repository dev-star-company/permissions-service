package auth_users_controller

import (
	"context"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/password"
	userSchema "permissions-service/internal/app/ent/user"
	"permissions-service/internal/infra/grpc_server/controllers"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *auth_users_proto.UpdateRequest) (*auth_users_proto.UpdateResponse, error) {
	if in.RequesterUuid == "" {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}
	defer tx.Rollback()

	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	userQ := tx.User.UpdateOneID(int(*in.Id)).SetUpdatedBy(requester.ID)

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

	if in.Password != nil && in.ConfirmPassword != nil && *in.Password == *in.ConfirmPassword {
		p, err := tx.Password.Create().
			SetPassword(*in.Password).
			SetCreatedBy(requester.ID).
			SetUpdatedBy(requester.ID).
			SetUserID(user.ID).
			Save(ctx)
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}

		oldPassword := tx.Password.Query().
			Where(password.IDNotIn(p.ID), password.HasUserWith(userSchema.IDEQ(user.ID))).
			Order(ent.Desc(password.FieldID)).
			FirstX(ctx)

		err = tx.Password.DeleteOneID(oldPassword.ID).Exec(ctx)
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}
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
