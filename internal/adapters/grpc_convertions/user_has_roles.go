package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/user_has_roles_proto"
)

func UserHasRolesToProto(user_has_roles *ent.UserHasRoles) *user_has_roles_proto.UserHasRoles {
	if user_has_roles == nil {
		return nil
	}

	e := user_has_roles_proto.UserHasRoles{
		RoleId: uint32(user_has_roles.RoleID),
		// UserId:    uint32(user_has_roles.UserID),
		CreatedAt: user_has_roles.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if user_has_roles.DeletedAt != nil {
		x := user_has_roles.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	return &e
}
