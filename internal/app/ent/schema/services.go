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
		field.String("name").NotEmpty(),
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
