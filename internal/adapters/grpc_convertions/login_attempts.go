package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/login_attempts_proto"
)

func LoginAttemptsToProto(login_attempts *ent.LoginAttempts) *login_attempts_proto.LoginAttempts {
	if login_attempts == nil {
		return nil
	}

	e := login_attempts_proto.LoginAttempts{
		Id:         uint32(login_attempts.ID),
		Successful: bool(login_attempts.Successful),
		CreatedAt:  login_attempts.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if login_attempts.DeletedAt != nil {
		x := login_attempts.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	return &e
}
