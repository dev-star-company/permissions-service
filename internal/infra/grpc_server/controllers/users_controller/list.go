package users_controller

import (
	"context"
	"fmt"
	"permission-service/generated_protos/users_proto"
	"permission-service/internal/adapters/grpc_controllers"
	"permission-service/internal/app/ent/schema"
	"permission-service/internal/app/ent/user"
)

func (c *controller) List(ct context.Context, in *users_proto.ListRequest) (*users_proto.ListResponse, error) {
	ctx := ct
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	// Get the user from the database
	query := tx.User.Query().
		Limit(int(in.Limit)).
		Offset(int(in.Offset))

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	if in.Name != nil {
		query = query.Where(user.NameContainsFold(*in.Name))
	}

	if in.Surname != nil {
		query = query.Where(user.SurnameContainsFold(*in.Surname))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying users: %w", err)
	}

	users, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying users: %w", err)
	}

	// Convert the users to the response format
	responseUsers := make([]*users_proto.User, len(users))
	for i, user := range users {
		responseUsers[i] = grpc_controllers.UserToProto(user)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	// Create and return the response
	return &users_proto.ListResponse{
		Rows:  responseUsers,
		Count: uint32(count),
	}, nil
}
