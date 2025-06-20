package schema

import (
	"context"
	"permissions-service/internal/app/ent/hook"
	"strings"

	gen "permissions-service/internal/app/ent"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.New()).Default(uuid.New).Unique(),
		field.String("name").NotEmpty(),
		field.String("surname").NotEmpty(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("emails", Email.Type),
		edge.To("passwords", Password.Type),
		edge.To("phones", Phone.Type),
		edge.To("ban", Ban.Type),
		edge.To("first_login", FirstLogin.Type),
		edge.To("login_attempts", LoginAttempts.Type),
		edge.To("roles", Role.Type).Through("user_has_roles", UserHasRoles.Type),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		// First hook.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					if name, exists := m.Name(); exists {
						m.SetName(strings.ToLower(name))
					}

					if surname, exists := m.Surname(); exists {
						m.SetSurname(strings.ToLower(surname))
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
