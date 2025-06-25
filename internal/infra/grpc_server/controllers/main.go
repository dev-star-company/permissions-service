package controllers

import (
	"context"
	"fmt"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/user"

	"github.com/google/uuid"
)

func GetRequesterId(tx *ent.Tx, ctx context.Context, requesterUuid string) (int, error) {
	uuidRequester, err := uuid.Parse(requesterUuid)
	if err != nil {
		return 0, fmt.Errorf("invalid requester UUID: %w", err)
	}
	requesterUser, err := tx.User.Query().
		Where(user.UUIDEQ(uuidRequester)).
		Only(ctx)
	if err != nil {
		return 0, fmt.Errorf("fetching requester user: %w", err)
	}
	if requesterUser == nil {
		return 0, fmt.Errorf("usuário solicitante não encontrado")
	}
	return requesterUser.ID, nil
}
