package roles_controller

import (
	"context"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/role"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *controller) Get(ctx context.Context, in *roles_proto.GetRequest) (*roles_proto.GetResponse, error) {

	if in.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "role ID is required")
	}

	if _, err := uuid.Parse(in.RequesterUuid); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid requester UUID")
	}

	roleEntity, err := c.Db.Role.
		Query().
		Where(role.IDEQ(int(in.Id))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "role not found")
		}
		return nil, status.Errorf(codes.Internal, "error retrieving role: %v", err)
	}

	return &roles_proto.GetResponse{
		Role: grpc_convertions.RoleToProto(roleEntity),
	}, nil
}
