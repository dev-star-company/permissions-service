package grpc_controllers

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"
)

func FirstLoginToProto(first_login *ent.FirstLogin) *first_login_proto.FirstLogin {
	if first_login == nil {
		return nil
	}

	e := first_login_proto.FirstLogin{
		Id:        uint32(first_login.ID),
		UserId:    uint32(*first_login.UserID),
		CreatedAt: first_login.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: first_login.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy: uint32(first_login.CreatedBy),
		UpdatedBy: uint32(first_login.UpdatedBy),
	}

	if first_login.DeletedAt != nil {
		x := first_login.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	if first_login.DeletedBy != nil {
		x := uint32(*first_login.DeletedBy)
		e.DeletedBy = &x
	}

	return &e
}
