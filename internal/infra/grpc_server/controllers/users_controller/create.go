package users_controller

import (
	"context"
	"fmt"
	"permission-service/generated_protos/users_proto"
	"permission-service/internal/adapters/grpc_controllers"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *users_proto.CreateRequest) (*users_proto.CreateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	// Create a new user in the database
	user, err := tx.User.Create().
		SetName(in.Name).
		SetSurname(in.Surname).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	email, err := tx.Email.Create().
		SetEmail(in.Email).
		SetUserID(user.ID).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	password, err := tx.Password.Create().
		SetPassword(in.Password).
		SetUserID(user.ID).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	phone, err := tx.Phone.Create().
		SetPhone(in.Phone).
		SetUserID(user.ID).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	user.Edges.Emails = append(user.Edges.Emails, email)
	user.Edges.Passwords = append(user.Edges.Passwords, password)
	user.Edges.Phones = append(user.Edges.Phones, phone)

	return &users_proto.CreateResponse{
		User: grpc_controllers.UserToProto(user),
	}, nil
}
