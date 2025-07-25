package app

import (
	"permissions-service/internal/app/ent"
	"permissions-service/internal/infra/grpc_server/controllers/auth_users_controller"
	"permissions-service/internal/infra/grpc_server/controllers/ban_controller"
	"permissions-service/internal/infra/grpc_server/controllers/first_login_controller"
	"permissions-service/internal/infra/grpc_server/controllers/login_attempts_controller"
	"permissions-service/internal/infra/grpc_server/controllers/permission_controller"
	"permissions-service/internal/infra/grpc_server/controllers/role_has_permissions_controller"
	"permissions-service/internal/infra/grpc_server/controllers/roles_controller"
	"permissions-service/internal/infra/grpc_server/controllers/service_controller"
	"permissions-service/internal/infra/grpc_server/controllers/user_has_roles_controller"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/ban_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/login_attempts_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/service_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/user_has_roles_proto"

	"google.golang.org/grpc"
)

func RegisterControllers(grpcServer *grpc.Server, client *ent.Client) {
	ban_proto.RegisterBanServiceServer(grpcServer, ban_controller.New(client))
	first_login_proto.RegisterFirstLoginServiceServer(grpcServer, first_login_controller.New(client))
	login_attempts_proto.RegisterLoginAttemptsServiceServer(grpcServer, login_attempts_controller.New(client))
	permission_proto.RegisterPermissionServiceServer(grpcServer, permission_controller.New(client))
	role_has_permissions_proto.RegisterRoleHasPermissionServiceServer(grpcServer, role_has_permissions_controller.New(client))
	roles_proto.RegisterRoleServiceServer(grpcServer, roles_controller.New(client))
	service_proto.RegisterServiceServiceServer(grpcServer, service_controller.New(client))
	auth_users_proto.RegisterAuthUsersServiceServer(grpcServer, auth_users_controller.New(client))
	user_has_roles_proto.RegisterUserHasRolesServiceServer(grpcServer, user_has_roles_controller.New(client))
}
