package grpc_controllers

import (
	"github.com/dev-star-company/protos-go/permission-service/generated_protos/service_proto"

	"permission-service/internal/app/ent"
)

func ServiceToProto(service *ent.Services) *service_proto.Service {
	if service == nil {
		return nil
	}

	cur := &service_proto.Service{
		Id:        uint32(service.ID),
		Name:      service.Name,
		CreatedBy: uint32(service.CreatedBy),
		UpdatedBy: uint32(service.UpdatedBy),
		CreatedAt: service.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: service.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if service.DeletedAt != nil {
		x := service.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if service.DeletedBy != nil {
		x := uint32(*service.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
