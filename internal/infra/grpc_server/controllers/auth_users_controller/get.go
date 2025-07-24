package auth_users_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/user"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *controller) Get(ctx context.Context, in *auth_users_proto.GetRequest) (*auth_users_proto.GetResponse, error) {

	// Retrieve the user from the database
	user, err := c.Db.User.
		Query().
		Where(user.IDEQ(int(in.Id))).
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

func GetRoleIDsByUserID(ctx context.Context, db *ent.Client, userID uint32) ([]uint32, error) {
	u, err := db.User.
		Query().
		Where(user.IDEQ(int(userID))).
		WithRoles().
		Only(ctx)

	if err != nil {
		return nil, err
	}

	roleIDs := make([]uint32, 0, len(u.Edges.Roles))
	for _, role := range u.Edges.Roles {
		roleIDs = append(roleIDs, uint32(role.ID))
	}

	return roleIDs, nil
}
