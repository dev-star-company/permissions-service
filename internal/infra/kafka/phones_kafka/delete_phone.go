package phones_kafka

import (
	"context"
	"permissions-service/internal/app/ent/phone"
	"permissions-service/internal/pkg/utils"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *phonesKafka) DeletePhone(u connection.SyncPhoneStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	_, err = tx.Phone.Delete().Where(phone.UUIDEQ(u.Uuid)).Exec(context.Background())
	if err != nil {
		return utils.Rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Successfully deleted phone, return nil error
	return nil
}
