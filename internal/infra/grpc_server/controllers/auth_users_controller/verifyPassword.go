package auth_users_controller

import (
	"context"
	"errors"
	"fmt"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/email"
	"permissions-service/internal/app/ent/phone"
	"permissions-service/internal/app/ent/user"
	"permissions-service/internal/pkg/utils/hash_password"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
	"github.com/dev-star-company/service-errors/errs"
	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TO DO: add soft ban after 3 failed attempts
func (c *controller) VerifyPassword(ctx context.Context, in *auth_users_proto.VerifyPasswordRequest) (*auth_users_proto.VerifyPasswordResponse, error) {
	if in.Password == "" {
		return nil, errs.BadRequest(errors.New("password is required"))
	}

	if in.Uuid == nil && in.Email == nil && in.Phone == nil {
		return nil, errs.BadRequest(errors.New("either uuid, email or phone is required"))
	}

	userQ := c.Db.User.Query()

	if in.Uuid != nil && *in.Uuid != "" {
		uuidRequester, err := uuid.Parse(*in.Uuid)
		if err != nil {
			return nil, fmt.Errorf("invalid requester UUID: %w", err)
		}
		userQ = userQ.Where(user.UUIDEQ(uuidRequester))
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
			return nil, errs.UserNotFound(0)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(u.Edges.Passwords) != 1 {
		// it's an internal error because if the user has no passwords, they should not be found in the first place
		// but we are handling it gracefully
		// and there is a bug in the CRUD process
		return nil, errs.InternalError(errors.New("user has no passwords or multiple passwords, this is an internal error"))
	}

	r := auth_users_proto.VerifyPasswordResponse{
		Success: false,
	}

	if hash_password.Check(in.Password, u.Edges.Passwords[0].Password) {
		r.Success = true
		r.User = grpc_convertions.UserToProto(u)
	}

	return &r, nil
}
