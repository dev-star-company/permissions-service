package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FirstLogin holds the schema definition for the FirstLogin entity.
type FirstLogin struct {
	ent.Schema
}

// Fields of the FirstLogin.
func (FirstLogin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Nillable().Unique(),
		field.Bool("successful").Default(false),
	}
}

func (FirstLogin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the FirstLogin.
func (FirstLogin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("first_login").Field("user_id").Unique().Required(),
	}
}
