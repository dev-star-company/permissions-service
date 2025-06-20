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
		UserId:     uint32(login_attempts.UserID),
		Successful: bool(login_attempts.Successful),
		CreatedAt:  login_attempts.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  login_attempts.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy:  uint32(login_attempts.CreatedBy),
		UpdatedBy:  uint32(login_attempts.UpdatedBy),
	}

	if login_attempts.DeletedAt != nil {
		x := login_attempts.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	if login_attempts.DeletedBy != nil {
		x := uint32(*login_attempts.DeletedBy)
		e.DeletedBy = &x
	}

	return &e
}
