package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("service_id"),
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional().NotEmpty(),
		field.Bool("is_active").Default(true),
		field.String("internal_name").NotEmpty().Unique(),
	}
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("permissions").Through("role_has_permissions", RoleHasPermissions.Type),
	}
}
