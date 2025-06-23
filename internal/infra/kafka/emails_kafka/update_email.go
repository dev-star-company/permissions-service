package emails_kafka

import (
	"context"
	"fmt"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/email"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils"
	"permissions-service/internal/pkg/utils/parser"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *emailsKafka) UpdateEmail(u connection.SyncEmailStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	query := tx.Email.Update().Where(email.UUIDEQ(u.Uuid))

	_, err = tx.User.Query().
		Where(
			user.UUIDEQ(u.UserUuid),
		).First(context.Background())

	if err != nil || !ent.IsNotFound(err) {
		fmt.Println("Error checking if user exists:", err)
	}

	if u.UpdatedBy != nil {
		query = query.SetUpdatedBy(int(*u.UpdatedBy))
	}
	if u.UpdatedAt != nil {
		query = query.SetUpdatedAt(*parser.ParseTimeTime(u.UpdatedAt))
	}
	_, err = query.Save(context.Background())

	if err != nil {
		return utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully updated email, return nil error
	return nil
}
