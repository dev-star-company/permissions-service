package phones_kafka

import (
	"context"
	"fmt"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/phone"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils"
	"permissions-service/internal/pkg/utils/parser"

	"github.com/dev-star-company/kafka-go/connection"
)

func (c *phonesKafka) UpdatePhone(u connection.SyncPhoneStruct) error {
	tx, err := c.db.Tx(context.Background())
	if err != nil {
		return err
	}

	query := tx.Phone.Update().Where(phone.UUIDEQ(u.Uuid))

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

	// Successfully updated phone, return nil error
	return nil
}
