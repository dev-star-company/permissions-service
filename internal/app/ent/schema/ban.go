package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ban holds the schema definition for the Ban entity.
type Ban struct {
	ent.Schema
}

// Fields of the Ban.
func (Ban) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Time("expires_at").Nillable(),
	}
}

func (Ban) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Ban.
func (Ban) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("ban").Field("user_id").Unique().Required(),
	}
}
