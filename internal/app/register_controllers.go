package app

import (
	"permissions-service/internal/app/ent"
	"permissions-service/internal/infra/grpc_server/controllers/ban_controller"
	"permissions-service/internal/infra/grpc_server/controllers/first_login_controller"
	"permissions-service/internal/infra/grpc_server/controllers/login_attempts_controller"
	"permissions-service/internal/infra/grpc_server/controllers/password_controller"
	"permissions-service/internal/infra/grpc_server/controllers/permission_controller"
	"permissions-service/internal/infra/grpc_server/controllers/role_has_permissions_controller"
	"permissions-service/internal/infra/grpc_server/controllers/roles_controller"
	"permissions-service/internal/infra/grpc_server/controllers/service_controller"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/ban_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/first_login_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/login_attempts_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/password_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/permission_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/role_has_permissions_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/roles_proto"
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/service_proto"

	"google.golang.org/grpc"
)

func RegisterControllers(grpcServer *grpc.Server, client *ent.Client) {
	ban_proto.RegisterServiceServer(grpcServer, ban_controller.New(client))
	first_login_proto.RegisterServiceServer(grpcServer, first_login_controller.New(client))
	login_attempts_proto.RegisterServiceServer(grpcServer, login_attempts_controller.New(client))
	permission_proto.RegisterServiceServer(grpcServer, permission_controller.New(client))
	role_has_permissions_proto.RegisterServiceServer(grpcServer, role_has_permissions_controller.New(client))
	roles_proto.RegisterServiceServer(grpcServer, roles_controller.New(client))
	service_proto.RegisterServiceServer(grpcServer, service_controller.New(client))
	password_proto.RegisterServiceServer(grpcServer, password_controller.New(client))
}
