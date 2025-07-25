package grpc_convertions

import (
	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/service_proto"

	"permissions-service/internal/app/ent"
)

func ServiceToProto(service *ent.Services) *service_proto.ServiceDto {
	if service == nil {
		return nil
	}

	cur := &service_proto.ServiceDto{
		Id:        uint32(service.ID),
		Name:      service.Name,
		CreatedAt: service.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if service.DeletedAt != nil {
		x := service.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
