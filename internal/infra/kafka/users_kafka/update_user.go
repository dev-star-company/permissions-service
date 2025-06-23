package users_kafka

import (
	"context"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *usersKafka) UpdateUser(u connection.SyncUserStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	query := tx.User.Update().Where(user.UUIDEQ(u.Uuid))

	if u.Name != nil {
		query = query.SetName(*u.Name)
	}
	if u.Surname != nil {
		query = query.SetSurname(*u.Surname)
	}
	if u.UpdatedBy != nil {
		query = query.SetUpdatedBy(int(*u.UpdatedBy))
	}
	if u.UpdatedAt != nil {
		query = query.SetUpdatedAt(*u.UpdatedAt)
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
