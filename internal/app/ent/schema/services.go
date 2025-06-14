package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Services holds the schema definition for the Services entity.
type Services struct {
	ent.Schema
}

// Fields of the Services.
func (Services) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional().NotEmpty(),
		field.String("internal_name").NotEmpty().Unique(),
		field.Bool("is_active").Default(true),
	}
}

func (Services) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Services.
func (Services) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permission", Permission.Type),
	}
}
