package users_kafka

import (
	"context"
	"fmt"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils"
	"permissions-service/internal/pkg/utils/parser"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *usersKafka) CreateUser(u connection.SyncUserStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	userExists, err := tx.User.Query().
		Where(
			user.UUIDEQ(u.Uuid),
		).First(context.Background())

	if err != nil {
		fmt.Println("Error checking if user exists:", err)
	}
	if userExists != nil {
		_, err = tx.User.
			UpdateOne(userExists).
			SetName(*u.Name).
			SetSurname(*u.Surname).
			SetUpdatedBy(int(*u.UpdatedBy)).
			SetUpdatedAt(*parser.ParseTimeTime(u.UpdatedAt)).
			Save(context.Background())
	} else {
		_, err = tx.User.Create().
			SetUUID(u.Uuid).
			SetName(*u.Name).
			SetSurname(*u.Surname).
			SetCreatedBy(int(*u.CreatedBy)).
			SetCreatedAt(*parser.ParseTimeTime(u.CreatedAt)).
			SetUpdatedBy(int(*u.UpdatedBy)).
			SetUpdatedAt(*parser.ParseTimeTime(u.UpdatedAt)).
			Save(context.Background())
	}

	if err != nil {
		return utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully created user, return nil error
	return nil
}
