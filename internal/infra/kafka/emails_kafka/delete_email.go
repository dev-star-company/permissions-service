package emails_kafka

import (
	"context"
	"permissions-service/internal/app/ent/email"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *emailsKafka) DeleteEmail(u connection.SyncEmailStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	_, err = tx.Email.Delete().Where(email.UUIDEQ(u.Uuid)).Exec(context.Background())
	if err != nil {
		return utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully deleted email, return nil error
	return nil
}
