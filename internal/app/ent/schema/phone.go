package schema

import (
	"context"
	"permissions-service/internal/app/ent/hook"
	"permissions-service/internal/app/ent/phone"

	gen "permissions-service/internal/app/ent"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Phone holds the schema definition for the Phone entity.
type Phone struct {
	ent.Schema
}

// Fields of the Phone.
func (Phone) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").NotEmpty(),
		field.Int("user_id"),
		field.Bool("main").Default(false),
	}
}

func (Phone) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Phone.
func (Phone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("phones").Field("user_id").Unique().Required(),
	}
}

func (Phone) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.PhoneFunc(func(ctx context.Context, m *gen.PhoneMutation) (ent.Value, error) {
					client := m.Client()
					userID, _ := m.UserID()
					exists, err := client.Phone.Query().
						Where(
							phone.UserID(userID),
							phone.Main(true),
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
