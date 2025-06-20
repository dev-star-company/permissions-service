package roles_controller

import (
	"context"
	"fmt"
	"permissions-service/internal/adapters/grpc_convertions"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/app/ent/role"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *controller) Get(ctx context.Context, in *roles_proto.GetRequest) (*roles_proto.GetResponse, error) {
	role, err := c.Db.Role.Query().
		Where(role.IDEQ(int(in.Id))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "role not found")
		}

		return nil, fmt.Errorf("querying role: %w", err)
	}

	return &roles_proto.GetResponse{
		Role: grpc_convertions.RoleToProto(role),
	}, nil
}
