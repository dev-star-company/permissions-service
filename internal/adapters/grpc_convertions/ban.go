package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/ban_proto"
)

func BanToProto(ban *ent.Ban) *ban_proto.Ban {
	if ban == nil {
		return nil
	}

	e := ban_proto.Ban{
		Id:        uint32(ban.ID),
		ExpiresAt: ban.ExpiresAt.Format("2006-01-02 15:04:05"),
		CreatedAt: ban.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if ban.DeletedAt != nil {
		x := ban.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	return &e
}
