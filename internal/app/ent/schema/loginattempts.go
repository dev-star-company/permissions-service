package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LoginAttempts holds the schema definition for the LoginAttempts entity.
type LoginAttempts struct {
	ent.Schema
}

// Fields of the LoginAttempts.
func (LoginAttempts) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Bool("successful").Default(false),
	}
}

func (LoginAttempts) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the LoginAttempts.
func (LoginAttempts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("login_attempts").Field("user_id").Unique().Required(),
	}
}
