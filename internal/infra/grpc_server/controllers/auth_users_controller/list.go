package auth_users_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent/email"
	"permissions-service/internal/app/ent/phone"
	"permissions-service/internal/app/ent/schema"
	"permissions-service/internal/app/ent/user"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
)

func (c *controller) List(ctx context.Context, in *auth_users_proto.ListRequest) (*auth_users_proto.ListResponse, error) {

	query := c.Db.User.Query()

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	if in.Name != nil {
		query = query.Where(user.NameContainsFold(*in.Name))
	}

	if in.Surname != nil {
		query = query.Where(user.SurnameContainsFold(*in.Surname))
	}

	if in.Phone != nil {
		query = query.Where(user.HasPhonesWith(phone.PhoneContainsFold(*in.Phone)))
	}

	if in.Email != nil {
		query = query.Where(user.HasEmailsWith(email.EmailContainsFold(*in.Email)))
	}

	if in.Relations != nil {
		if in.Relations.Emails {
			query = query.WithEmails()
		}

		if in.Relations.Phones {
			query = query.WithPhones()
		}

		if in.Relations.Roles {
			query = query.WithRoles()
		}
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying users: %w", err)
	}

	if in.Limit != nil && *in.Limit > 0 {
		query = query.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		query = query.Offset(int(*in.Offset))
	}

	users, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying users: %w", err)
	}

	// Convert the users to the response format
	responseUsers := make([]*auth_users_proto.User, len(users))
	for i, user := range users {
		responseUsers[i] = grpc_convertions.UserToProto(user)
	}

	// Create and return the response
	return &auth_users_proto.ListResponse{
		Rows:  responseUsers,
		Count: uint32(count),
	}, nil
}
