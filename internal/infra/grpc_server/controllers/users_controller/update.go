package users_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/adapters/grpc_controllers"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/password"
	"permissions-service/internal/pkg/utils"
	"permissions-service/internal/pkg/utils/hash_password"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/users_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) UpdateUser(ctx context.Context, in *users_proto.UpdateRequest) (*users_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	// Update the user in the database
	userQ := tx.User.UpdateOneID(int(in.Id))

	if in.Name != nil {
		userQ.SetName(*in.Name)
	}

	if in.Surname != nil {
		userQ.SetSurname(*in.Surname)
	}

	if in.Password != nil && in.ConfirmPassword != nil && *in.Password == *in.ConfirmPassword {
		hashedPassword, _ := hash_password.Hash(*in.Password)
		p, err := tx.Password.Create().
			SetPassword(hashedPassword).
			SetCreatedBy(int(in.RequesterId)).
			SetUpdatedBy(int(in.RequesterId)).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		np := tx.Password.Query().Where(password.IDNotIn(p.ID)).Order(ent.Asc(password.FieldID)).FirstX(ctx)
		err = tx.Password.DeleteOneID(np.ID).Exec(ctx)
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}
	}

	user, err := userQ.Save(ctx)
	if err != nil {
		return nil, err
	}

	// Create and return the response
	return &users_proto.UpdateResponse{
		User: grpc_controllers.UserToProto(user),
	}, nil

}
