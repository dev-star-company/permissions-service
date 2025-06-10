package schema

import (
	"context"
	"permissions-service/internal/app/ent/email"
	"permissions-service/internal/app/ent/hook"
	"strings"

	gen "permissions-service/internal/app/ent"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Email holds the schema definition for the Email entity.
type Email struct {
	ent.Schema
}

// Fields of the Email.
func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().Unique(),
		field.Int("user_id"),
		field.Bool("main").Default(false),
	}
}

func (Email) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Email.
func (Email) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("emails").Field("user_id").Unique().Required(),
	}
}

func (Email) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.EmailFunc(func(ctx context.Context, m *gen.EmailMutation) (ent.Value, error) {
					if email, exists := m.Email(); exists {
						m.SetEmail(strings.ToLower(email))
					}

					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.EmailFunc(func(ctx context.Context, m *gen.EmailMutation) (ent.Value, error) {
					client := m.Client()
					userID, _ := m.UserID()
					exists, err := client.Email.Query().
						Where(
							email.UserID(userID),
							email.Main(true),
						).Exist(ctx)
					if err != nil {
						return nil, err
					}
					if exists {
						// Skip marking this email as main since another one already is
						return next.Mutate(ctx, m)
					}
					m.SetMain(true)
					return next.Mutate(ctx, m)

				})
			},
			ent.OpCreate,
		),
	}
}
