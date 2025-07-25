package grpc_convertions

import (
	"permissions-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/permissions_service/generated_protos/auth_users_proto"
)

func UserToProto(user *ent.User) *auth_users_proto.User {
	if user == nil {
		return nil
	}

	u := auth_users_proto.User{
		Id:        uint32(user.ID),
		Surname:   user.Surname,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if len(user.Edges.Emails) != 0 {
		u.Emails = make([]*auth_users_proto.Email, len(user.Edges.Emails))
		for i, email := range user.Edges.Emails {
			u.Emails[i] = EmailToProto(email)
		}
	}

	if len(user.Edges.Phones) != 0 {
		u.Phones = make([]*auth_users_proto.Phone, len(user.Edges.Phones))
		for i, phone := range user.Edges.Phones {
			u.Phones[i] = PhoneToProto(phone)
		}
	}

	if user.DeletedAt != nil {
		x := user.DeletedAt.Format("2006-01-02 15:04:05")
		u.DeletedAt = &x
	}

	return &u
}
