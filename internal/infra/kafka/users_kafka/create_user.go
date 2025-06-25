package users_kafka

import (
	"context"
	"fmt"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils"

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

	if err != nil || !ent.IsNotFound(err) {
		fmt.Println("Error checking if user exists:", err)
	}
	if userExists != nil {
		_, err = tx.User.
			UpdateOne(userExists).
			SetName(*u.Name).
			SetSurname(*u.Surname).
			SetUpdatedBy(int(*u.UpdatedBy)).
			SetUpdatedAt(*u.UpdatedAt).
			Save(context.Background())
	} else {
		_, err = tx.User.Create().
			SetUUID(u.Uuid).
			SetName(*u.Name).
			SetSurname(*u.Surname).
			SetCreatedBy(int(*u.CreatedBy)).
			SetCreatedAt(*u.CreatedAt).
			SetUpdatedBy(int(*u.UpdatedBy)).
			SetUpdatedAt(*u.UpdatedAt).
			Save(context.Background())
	}

	if err != nil {
		return utils.Rollback(tx, err)
	}

	for _, phone := range u.Phones {
		newPhone, err := tx.Phone.Create().
			SetUUID(phone.Uuid).
			SetPhone(*phone.Phone).
			SetMain(*phone.Main).
			SetUserID(userExists.ID).
			SetCreatedBy(int(*phone.CreatedBy)).
			SetCreatedAt(*phone.CreatedAt).
			SetUpdatedBy(int(*phone.UpdatedBy)).
			SetUpdatedAt(*phone.UpdatedAt).
			Save(context.Background())
		if err != nil {
			return utils.Rollback(tx, fmt.Errorf("error creating phone: %w", err))
		}
		fmt.Printf("Created phone: %v\n", newPhone)
	}

	for _, email := range u.Emails {
		newEmail, err := tx.Email.Create().
			SetUUID(email.Uuid).
			SetEmail(*email.Email).
			SetMain(*email.Main).
			SetUserID(userExists.ID).
			SetCreatedBy(int(*email.CreatedBy)).
			SetCreatedAt(*email.CreatedAt).
			SetUpdatedBy(int(*email.UpdatedBy)).
			SetUpdatedAt(*email.UpdatedAt).
			Save(context.Background())
		if err != nil {
			return utils.Rollback(tx, fmt.Errorf("error creating email: %w", err))
		}
		fmt.Printf("Created email: %v\n", newEmail)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully created user, return nil error
	return nil
}
