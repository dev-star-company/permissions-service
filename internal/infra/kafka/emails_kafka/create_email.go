package emails_kafka

import (
	"context"
	"fmt"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/email"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *emailsKafka) CreateEmail(u connection.SyncEmailStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	emailExists, err := tx.Email.Query().
		Where(
			email.UUIDEQ(u.Uuid),
		).First(context.Background())

	if err != nil {
		fmt.Println("Error checking if email exists:", err)
	}

	_, err = tx.User.Query().
		Where(
			user.UUIDEQ(u.UserUuid),
		).First(context.Background())

	if err != nil || !ent.IsNotFound(err) {
		fmt.Println("Error checking if user exists:", err)
	}

	if emailExists != nil {
		_, err = tx.Email.
			UpdateOne(emailExists).
			SetUpdatedBy(int(*u.UpdatedBy)).
			SetUpdatedAt(*u.UpdatedAt).
			Save(context.Background())
	} else {
		_, err = tx.Email.Create().
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

	// Successfully created email, return nil error
	return nil
}
