package grpc_controllers

import (
	"permission-service/generated_protos/ban_proto"
	"permission-service/internal/app/ent"
)

func BanToProto(ban *ent.Ban) *ban_proto.Ban {
	if ban == nil {
		return nil
	}

	e := ban_proto.Ban{
		Id:        uint32(ban.ID),
		UserId:    uint32(ban.UserID),
		ExpiresAt: ban.ExpiresAt.Format("2006-01-02 15:04:05"),
		CreatedAt: ban.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: ban.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy: uint32(ban.CreatedBy),
		UpdatedBy: uint32(ban.UpdatedBy),
	}

	if ban.DeletedAt != nil {
		x := ban.DeletedAt.Format("2006-01-02 15:04:05")
		e.DeletedAt = &x
	}

	if ban.DeletedBy != nil {
		x := uint32(*ban.DeletedBy)
		e.DeletedBy = &x
	}

	return &e
}
