package users_kafka

import (
	"context"
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

type UserKafka struct {
	db *ent.Client
}

func (c *UserKafka) SyncUser(user connection.SyncUserStruct) {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		// Handle transaction start error
		return
	}

	_, err = tx.User.Create().
		SetId(user.ID).
		SetName(*user.Name).
		SetSurname(*user.Surname).
		SetCreatedBy(int(*user.CreatedBy)).
		SetCreatedAt(*user.CreatedAt).
		SetUpdatedBy(int(*user.UpdatedBy)).
		SetUpdatedAt(*user.UpdatedAt).
		Save(context.Background())

	if err != nil {
		// Handle create error
		return
	}

	if err := tx.Commit(); err != nil {
		// Handle commit error
		return
	}

	// Successfully created user

}
