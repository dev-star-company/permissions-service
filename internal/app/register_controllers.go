package app

import (
	"permission-service/internal/app/ent"
	"permission-service/internal/infra/grpc_server/controllers/ban_controller"
	"permission-service/internal/infra/grpc_server/controllers/first_login_controller"
	"permission-service/internal/infra/grpc_server/controllers/login_attempts_controller"
	"permission-service/internal/infra/grpc_server/controllers/permission_controller"
	"permission-service/internal/infra/grpc_server/controllers/role_has_permissions_controller"
	"permission-service/internal/infra/grpc_server/controllers/roles_controller"
	"permission-service/internal/infra/grpc_server/controllers/service_controller"
	"permission-service/internal/infra/grpc_server/controllers/users_controller"

	"github.com/dev-star-company/protos-go/permission-service/generated_protos/ban_proto"
	"github.com/dev-star-company/protos-go/permission-service/generated_protos/first_login_proto"
	"github.com/dev-star-company/protos-go/permission-service/generated_protos/login_attempts_proto"
	"github.com/dev-star-company/protos-go/permission-service/generated_protos/permission_proto"
	"github.com/dev-star-company/protos-go/permission-service/generated_protos/role_has_permissions_proto"
	"github.com/dev-star-company/protos-go/permission-service/generated_protos/roles_proto"
	"github.com/dev-star-company/protos-go/permission-service/generated_protos/service_proto"
	"github.com/dev-star-company/protos-go/permission-service/generated_protos/users_proto"

	"google.golang.org/grpc"
)

func RegisterControllers(grpcServer *grpc.Server, client *ent.Client) {
	ban_proto.RegisterBanServiceServer(grpcServer, ban_controller.New(client))
	first_login_proto.RegisterFirstLoginServiceServer(grpcServer, first_login_controller.New(client))
	login_attempts_proto.RegisterLoginAttemptsServiceServer(grpcServer, login_attempts_controller.New(client))
	permission_proto.RegisterPermissionServiceServer(grpcServer, permission_controller.New(client))
	role_has_permissions_proto.RegisterRoleHasPermissionsServiceServer(grpcServer, role_has_permissions_controller.New(client))
	roles_proto.RegisterRolesServiceServer(grpcServer, roles_controller.New(client))
	service_proto.RegisterServiceServiceServer(grpcServer, service_controller.New(client))
	users_proto.RegisterUsersServiceServer(grpcServer, users_controller.New(client))
}
