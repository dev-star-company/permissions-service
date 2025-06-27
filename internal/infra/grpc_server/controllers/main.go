package controllers

import (
	"context"
	"fmt"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils/parser"

	"github.com/dev-star-company/service-errors/errs"
)

func GetUserFromUuid(tx *ent.Tx, ctx context.Context, requesterUuid string) (*ent.User, error) {
	if requesterUuid == "" {
		return nil, errs.RequesterIDRequired()
	}

	uuidRequester, err := parser.Uuid(requesterUuid)
	if err != nil {
		return nil, fmt.Errorf("invalid requester UUID: %w", err)
	}

	requesterUser, err := tx.User.Query().
		Where(user.UUIDEQ(uuidRequester)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching requester user: %w", err)
	}
	if requesterUser == nil {
		return nil, errs.UserNotFound(0)
	}
	return requesterUser, nil
}
