package phones_kafka

import (
	"context"
	"fmt"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/phone"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *phonesKafka) CreatePhone(u connection.SyncPhoneStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	phoneExists, err := tx.Phone.Query().
		Where(
			phone.UUIDEQ(u.Uuid),
		).First(context.Background())

	if err != nil {
		fmt.Println("Error checking if phone exists:", err)
	}

	_, err = tx.User.Query().
		Where(
			user.UUIDEQ(u.UserUuid),
		).First(context.Background())

	if err != nil || !ent.IsNotFound(err) {
		fmt.Println("Error checking if user exists:", err)
	}

	if phoneExists != nil {
		_, err = tx.Phone.
			UpdateOne(phoneExists).
			SetUpdatedBy(int(*u.UpdatedBy)).
			SetUpdatedAt(*u.UpdatedAt).
			Save(context.Background())
	} else {
		_, err = tx.Phone.Create().
			SetUUID(u.Uuid).
			SetCreatedBy(int(*u.CreatedBy)).
			SetCreatedAt(*u.CreatedAt).
			SetUpdatedBy(int(*u.UpdatedBy)).
			SetUpdatedAt(*u.UpdatedAt).
			Save(context.Background())
	}

	if err != nil {
		return utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully created phone, return nil error
	return nil
}
