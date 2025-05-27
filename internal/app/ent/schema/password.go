package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Password holds the schema definition for the Password entity.
type Password struct {
	ent.Schema
}

// Fields of the Password.
func (Password) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").NotEmpty(),
		field.Int("user_id"),
	}
}

func (Password) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Password.
func (Password) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("passwords").Field("user_id").Unique().Required(),
	}
}

// func (Password) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		// First hook.
// 		hook.On(
// 			func(next ent.Mutator) ent.Mutator {
// 				return hook.PasswordFunc(func(ctx context.Context, m *gen.PasswordMutation) (ent.Value, error) {
// 					if password, exists := m.Password(); exists {
// 						hashed_pw, err := hash_password.Hash(password)
// 						if err != nil {
// 							return nil, err
// 						}
// 						m.SetPassword(hashed_pw)
// 					}

// 					return next.Mutate(ctx, m)
// 				})
// 			},
// 			// Limit the hook only for these operations.
// 			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
// 		),
// 	}
// }
