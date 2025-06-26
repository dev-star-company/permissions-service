package parser

import (
	"fmt"

	"github.com/google/uuid"
)

func Uuid(requesterUuid string) (uuid.UUID, error) {
	uuidRequester, err := uuid.Parse(requesterUuid)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid requester UUID: %w", err)
	}
	return uuidRequester, nil
}
