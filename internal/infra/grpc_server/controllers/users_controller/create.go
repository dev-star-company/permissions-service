package users_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/adapters/kafka_dtos"
	"permissions-service/internal/app/ent/email"
	"permissions-service/internal/app/ent/phone"

	"github.com/dev-star-company/kafka-go/actions"
	"github.com/dev-star-company/kafka-go/connection"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *auth_users_proto.CreateRequest) (*auth_users_proto.CreateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.RequesterIDRequired()
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

	existsEmail, err := c.Db.Email.Query().Where(email.EmailEQ(in.Email)).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("verificando email existente: %w", err)
	}
	if existsEmail {
		return nil, fmt.Errorf("email já cadastrado")
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

	existsPhone, err := c.Db.Phone.Query().Where(phone.PhoneEQ(in.Phone)).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("verificando telefone existente: %w", err)
	}
	if existsPhone {
		return nil, fmt.Errorf("telefone já cadastrado")
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

	userToSync := connection.Message[connection.SyncUserStruct]{
		Action:  actions.CREATE,
		Payload: kafka_dtos.ToKafkaUser(*user),
	}
	c.k.PublishToSyncUsers(userToSync)

	emailToSync := connection.Message[connection.SyncEmailStruct]{
		Action:  actions.CREATE,
		Payload: kafka_dtos.ToKafkaEmail(*email),
	}
	c.k.PublishToSyncEmails(emailToSync)

	phoneToSync := connection.Message[connection.SyncPhoneStruct]{
		Action:  actions.CREATE,
		Payload: kafka_dtos.ToKafkaPhone(*phone),
	}
	c.k.PublishToSyncPhones(phoneToSync)

	return &auth_users_proto.CreateResponse{
		User: grpc_convertions.UserToProto(user),
	}, nil
}
