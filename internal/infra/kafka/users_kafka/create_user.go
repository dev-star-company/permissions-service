package users_kafka

import (
	"context"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *usersKafka) CreateUser(user connection.SyncUserStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	_, err = tx.User.Create().
		SetName(*user.Name).
		SetSurname(*user.Surname).
		SetCreatedBy(int(*user.CreatedBy)).
		SetCreatedAt(*user.CreatedAt).
		SetUpdatedBy(int(*user.UpdatedBy)).
		SetUpdatedAt(*user.UpdatedAt).
		Save(context.Background())

	if err != nil {
		return utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully created user, return nil error
	return nil
}
