package grpc_controllers

import (
	"permission-service/generated_protos/users_proto"
	"permission-service/internal/app/ent"
)

func UserToProto(user *ent.User) *users_proto.User {
	if user == nil {
		return nil
	}

	u := users_proto.User{
		Id:        uint32(user.ID),
		Surname:   user.Surname,
		Name:      user.Name,
		CreatedBy: uint32(user.CreatedBy),
		UpdatedBy: uint32(user.UpdatedBy),
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if len(user.Edges.Emails) != 0 {
		u.Emails = make([]*users_proto.Email, len(user.Edges.Emails))
		for i, email := range user.Edges.Emails {
			u.Emails[i] = EmailToProto(email)
		}
	}

	if len(user.Edges.Phones) != 0 {
		u.Phones = make([]*users_proto.Phone, len(user.Edges.Phones))
		for i, phone := range user.Edges.Phones {
			u.Phones[i] = PhoneToProto(phone)
		}
	}

	if user.DeletedAt != nil {
		x := user.DeletedAt.Format("2006-01-02 15:04:05")
		u.DeletedAt = &x
	}

	if user.DeletedBy != nil {
		x := uint32(*user.DeletedBy)
		u.DeletedBy = &x
	}

	return &u
}
