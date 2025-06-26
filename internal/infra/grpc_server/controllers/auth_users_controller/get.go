package auth_users_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils/parser"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *controller) Get(ctx context.Context, in *auth_users_proto.GetRequest) (*auth_users_proto.GetResponse, error) {

	userUuid, err := parser.Uuid(in.Uuid)
	if err != nil {
		return nil, err
	}
	// Retrieve the user from the database
	user, err := c.Db.User.
		Query().
		Where(user.UUID(userUuid)).
		WithPhones().
		WithPasswords().
		WithEmails().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		return nil, fmt.Errorf("retrieving user: %w", err)
	}

	// Create and return the response
	return &auth_users_proto.GetResponse{
		User: grpc_convertions.UserToProto(user),
	}, nil

}
