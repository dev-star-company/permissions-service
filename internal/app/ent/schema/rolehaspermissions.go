package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RoleHasPermissions holds the schema definition for the RoleHasPermissions entity.
type RoleHasPermissions struct {
	ent.Schema
}

func (RoleHasPermissions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("role_id", "permission_id"),
	}
}

func (RoleHasPermissions) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the RoleHasPermissions.
func (RoleHasPermissions) Fields() []ent.Field {
	return []ent.Field{
		field.Int("role_id").Positive(),
		field.Int("permission_id").Positive(),
	}
}

// Edges of the RoleHasPermissions.
func (RoleHasPermissions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).Field("role_id").Unique().Required(),
		edge.To("permissions", Permission.Type).Field("permission_id").Unique().Required(),
	}
}

func (RoleHasPermissions) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "permission_id").Unique(),
	}
}
