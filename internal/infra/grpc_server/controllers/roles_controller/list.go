package roles_controller

import (
	"context"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent/role"
	"permissions-service/internal/app/ent/schema"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
)

func (c *controller) List(ctx context.Context, in *roles_proto.ListRequest) (*roles_proto.ListResponse, error) {

	rolesQ := c.Db.Role.Query()

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	if in.Name != nil {
		rolesQ = rolesQ.Where(role.NameContainsFold(*in.Name))
	}

	if in.Description != nil {
		rolesQ = rolesQ.Where(role.DescriptionContainsFold(*in.Description))
	}

	if in.IsActive != nil {
		rolesQ = rolesQ.Where(role.IsActiveEQ(*in.IsActive))
	}

	if in.Limit != nil && *in.Limit > 0 {
		rolesQ = rolesQ.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		rolesQ = rolesQ.Offset(int(*in.Offset))
	}

	// if in.OrderBy != nil {
	// 	if val := in.OrderBy.Id; val != nil {
	// 		if *val == "asc" {
	// 			rolesQ = rolesQ.Order(role.ByID())
	// 		} else if *val == "desc" {
	// 			rolesQ = rolesQ.Order(role.ByIDDesc())
	// 		}
	// 	}
	// 	if val := in.OrderBy.Name; val != nil {
	// 		if *val == "asc" {
	// 			rolesQ = rolesQ.Order(role.ByName())
	// 		} else if *val == "desc" {
	// 			rolesQ = rolesQ.Order(role.ByNameDesc())
	// 		}
	// 	}
	// 	if val := in.OrderBy.CreatedAt; val != nil {
	// 		if *val == "asc" {
	// 			rolesQ = rolesQ.Order(role.ByCreatedAt())
	// 		} else if *val == "desc" {
	// 			rolesQ = rolesQ.Order(role.ByCreatedAtDesc())
	// 		}
	// 	}
	// }

	roles, err := rolesQ.All(ctx)
	if err != nil {
		return nil, err
	}

	count, err := rolesQ.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	response := &roles_proto.ListResponse{
		Rows:  make([]*roles_proto.Role, len(roles)),
		Count: uint32(count),
	}

	for i, role := range roles {
		response.Rows[i] = grpc_convertions.RoleToProto(role)
	}

	return response, nil
}
