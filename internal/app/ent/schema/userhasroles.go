package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserHasRoles holds the schema definition for the UserHasRoles entity.
type UserHasRoles struct {
	ent.Schema
}

func (UserHasRoles) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (UserHasRoles) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "role_id"),
	}
}

// Fields of the UserHasRoles.
func (UserHasRoles) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Positive(),
		field.Int("role_id").Positive(),
	}
}

// Edges of the UserHasRoles.
func (UserHasRoles) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Field("user_id").Unique().Required(),
		edge.To("roles", Role.Type).Field("role_id").Unique().Required(),
	}
}

func (UserHasRoles) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "role_id").Unique(),
	}
}
