package user_has_roles_controller

import (
	"context"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/userhasroles"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/user_has_roles_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *user_has_roles_proto.GetRequest) (*user_has_roles_proto.GetResponse, error) {
	_, err := c.Db.UserHasRoles.
		Query().
		Where(userhasroles.RoleID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.UserHasRolesNotFound(int(in.Id))
	}

	return &user_has_roles_proto.GetResponse{
		RequesterUuid: in.RequesterUuid,
	}, nil
}
