package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional().NotEmpty(),
		field.Bool("is_active").Default(true),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type).Through("role_has_permissions", RoleHasPermissions.Type),
		edge.From("permission", User.Type).Ref("roles").Through("user_has_roles", UserHasRoles.Type),
	}
}

// func (Role) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		// First hook.
// 		hook.On(
// 			func(next ent.Mutator) ent.Mutator {
// 				return hook.RoleFunc(func(ctx context.Context, m *gen.RoleMutation) (ent.Value, error) {
// 					if name, exists := m.Name(); exists {
// 						m.SetName(strings.ToLower(name))
// 					}

// 					return next.Mutate(ctx, m)
// 				})
// 			},
// 			// Limit the hook only for these operations.
// 			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
// 		),
// 	}
// }
