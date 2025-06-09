package users_controller

import (
	"context"
	"permissions-service/internal/adapters/grpc_controllers"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/email"
	"permissions-service/internal/app/ent/phone"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils/hash_password"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/users_proto"

	"github.com/dev-star-company/service-errors/errs"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TO DO: add soft ban after 3 failed attempts
func (c *controller) VerifyPassword(ctx context.Context, in *users_proto.VerifyPasswordRequest) (*users_proto.VerifyPasswordResponse, error) {
	if in.RequesterId == 0 {
		return nil, status.Error(codes.InvalidArgument, errs.RequesterIdRequired().Error())
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if in.Id == nil && in.Email == nil && in.Phone == nil {
		return nil, status.Error(codes.InvalidArgument, "either id, email or phone is required")
	}

	userQ := c.Db.User.Query()

	if in.Id != nil && *in.Id != 0 {
		userQ = userQ.Where(user.ID(int(*in.Id)))
	}

	if in.Email != nil && *in.Email != "" {
		userQ = userQ.Where(user.HasEmailsWith(email.Email(*in.Email)))
	}

	if in.Phone != nil && *in.Phone != "" {
		userQ = userQ.Where(user.HasPhonesWith(phone.Phone(*in.Phone)))
	}

	userQ = userQ.WithPasswords()
	u, err := userQ.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.UserNotFound(int(*in.Id))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(u.Edges.Passwords) != 1 {
		// it's an internal error because if the user has no passwords, they should not be found in the first place
		// but we are handling it gracefully
		// and there is a bug in the CRUD process
		return nil, status.Error(codes.Internal, "internal error on passwords storage")
	}

	r := users_proto.VerifyPasswordResponse{
		Success: false,
	}

	if hash_password.Check(in.Password, u.Edges.Passwords[0].Password) {
		r.Success = true
		r.User = grpc_controllers.UserToProto(u)
	}

	return &r, nil
}
