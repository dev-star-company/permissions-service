package users_kafka

import (
	"context"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *usersKafka) UpdateUser(user connection.SyncUserStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	query := tx.User.UpdateOneID(int(user.ID))

	if user.Name != nil {
		query = query.SetName(*user.Name)
	}
	if user.Surname != nil {
		query = query.SetSurname(*user.Surname)
	}
	if user.UpdatedBy != nil {
		query = query.SetUpdatedBy(int(*user.UpdatedBy))
	}
	if user.UpdatedAt != nil {
		query = query.SetUpdatedAt(*user.UpdatedAt)
	}
	_, err = query.Save(context.Background())

	if err != nil {
		return utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully updated user, return nil error
	return nil
}
