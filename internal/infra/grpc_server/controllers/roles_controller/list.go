package roles_controller

import (
	"context"
	"errors"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/role"
	"permissions-service/internal/app/ent/schema"
	"strings"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
	"github.com/dev-star-company/service-errors/errs"
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
			if sort == "asc" {
				rolesQ = rolesQ.Order(ent.Asc("id"))
			} else if sort == "desc" {
				rolesQ = rolesQ.Order(ent.Desc("id"))
			}
		}

		if in.OrderBy.Name != nil {
			sort := strings.ToLower(*in.OrderBy.Name)
			if !validSorts[sort] {
				return nil, errs.InvalidArgument(errors.New("invalid sort for name"))
			}
			if sort == "asc" {
				rolesQ = rolesQ.Order(ent.Asc("name"))
			} else if sort == "desc" {
				rolesQ = rolesQ.Order(ent.Desc("name"))
			}
		}

		if in.OrderBy.CreatedAt != nil {
			sort := strings.ToLower(*in.OrderBy.CreatedAt)
			if !validSorts[sort] {
				return nil, errs.InvalidArgument(errors.New("invalid sort for created_at"))
			}
			if sort == "asc" {
				rolesQ = rolesQ.Order(ent.Asc("created_at"))
			} else if sort == "desc" {
				rolesQ = rolesQ.Order(ent.Desc("created_at"))
			}
		}
	}

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
