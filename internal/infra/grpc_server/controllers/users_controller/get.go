package users_controller

import (
	"context"
	"fmt"
	"permission-service/generated_protos/users_proto"
	"permission-service/internal/adapters/grpc_controllers"
	"permission-service/internal/app/ent"
	"permission-service/internal/app/ent/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *controller) Get(ctx context.Context, in *users_proto.GetRequest) (*users_proto.GetResponse, error) {
	// Retrieve the user from the database
	user, err := c.Db.User.
		Query().
		Where(user.ID(int(in.Id))).
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
	return &users_proto.GetResponse{
		User: grpc_controllers.UserToProto(user),
	}, nil

}
