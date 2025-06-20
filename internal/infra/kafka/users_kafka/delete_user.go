package users_kafka

import (
	"context"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *usersKafka) DeleteUser(user connection.SyncUserStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	err = tx.User.DeleteOneID(int(user.ID)).Exec(context.Background())
	if err != nil {
		return utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully deleted user, return nil error
	return nil
}
