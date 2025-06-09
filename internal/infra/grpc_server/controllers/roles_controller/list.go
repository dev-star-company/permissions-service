package roles_controller

import (
	"context"
	"permission-service/internal/adapters/grpc_controllers"
	"permission-service/internal/app/ent/schema"

	"github.com/dev-star-company/protos-go/permission-service/generated_protos/roles_proto"
)

func (c *controller) List(ct context.Context, in *roles_proto.ListRequest) (*roles_proto.ListResponse, error) {
	ctx := ct

	// Create a new role in the database
	rolesQ := c.Db.Role.Query()

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	roles, err := rolesQ.All(ctx)
	if err != nil {
		return nil, err
	}

	count, err := rolesQ.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := &roles_proto.ListResponse{
		Rows:  make([]*roles_proto.Role, len(roles)),
		Count: uint32(count),
	}

	for i, role := range roles {
		response.Rows[i] = grpc_controllers.RoleToProto(role)
	}

	return response, nil
}
