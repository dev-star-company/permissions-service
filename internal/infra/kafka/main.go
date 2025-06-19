package kafka

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/kafka-go/connection"
)

type kafka struct {
	Db *ent.Client
	c  *connection.Connectioner
}

func New(Db *ent.Client, c *connection.Connectioner) *kafka {
	return &kafka{
		Db: Db,
		c:  c,
	}
}
