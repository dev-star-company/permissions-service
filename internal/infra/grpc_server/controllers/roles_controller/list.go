package roles_controller

import (
	"context"
	"errors"
	"fmt"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/role"
	"permissions-service/internal/app/ent/schema"
	"strings"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *roles_proto.ListRequest) (*roles_proto.ListResponse, error) {

	// Get the role from the database
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

	validSorts := map[string]bool{
		"":     true,
		"asc":  true,
		"desc": true,
	}

	if in.OrderBy != nil {
		if in.OrderBy.Id != nil {
			sort := strings.ToLower(*in.OrderBy.Id)
			if !validSorts[sort] {
				return nil, errs.InvalidArgument(errors.New("invalid sort for id"))
			}
			switch sort {
			case "asc":
				rolesQ = rolesQ.Order(ent.Asc("id"))
			case "desc":
				rolesQ = rolesQ.Order(ent.Desc("id"))
			}
		}

		if in.OrderBy.Name != nil {
			sort := strings.ToLower(*in.OrderBy.Name)
			if !validSorts[sort] {
				return nil, errs.InvalidArgument(errors.New("invalid sort for name"))
			}
			switch sort {
			case "asc":
				rolesQ = rolesQ.Order(ent.Asc("name"))
			case "desc":
				rolesQ = rolesQ.Order(ent.Desc("name"))
			}
		}

		if in.OrderBy.CreatedAt != nil {
			sort := strings.ToLower(*in.OrderBy.CreatedAt)
			if !validSorts[sort] {
				return nil, errs.InvalidArgument(errors.New("invalid sort for created_at"))
			}
			switch sort {
			case "asc":
				rolesQ = rolesQ.Order(ent.Asc("created_at"))
			case "desc":
				rolesQ = rolesQ.Order(ent.Desc("created_at"))
			}
		}
	}

	count, err := rolesQ.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying roles: %w", err)
	}

	if in.Limit != nil && *in.Limit > 0 {
		rolesQ = rolesQ.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		rolesQ = rolesQ.Offset(int(*in.Offset))
	}

	roles, err := rolesQ.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying roles: %w", err)
	}

	responseRoles := make([]*roles_proto.Role, len(roles))
	for i, role := range roles {
		responseRoles[i] = grpc_convertions.RoleToProto(role)
	}

	return &roles_proto.ListResponse{
		Rows:  responseRoles,
		Count: uint32(count),
	}, nil
}
