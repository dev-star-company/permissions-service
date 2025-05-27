package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin holds the schema definition for the BaseMixin entity.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
		field.Int("created_by").Positive().Immutable(),
		field.Int("updated_by").Positive(),
		field.Int("deleted_by").Optional().Nillable(),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the BaseMixin.
// func (d BaseMixin) Interceptors() []ent.Interceptor {
// 	return []ent.Interceptor{
// 		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
// 			// Skip soft-delete, means include soft-deleted entities.
// 			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
// 				return nil
// 			}
// 			d.P(q)
// 			return nil
// 		}),
// 	}
// }

func (m BaseMixin) EdgesMixin(schemaName string) []ent.Edge {
	switch schemaName {
	case "Products":
		return []ent.Edge{
			edge.From("created_by_user", User.Type).
				Ref("created_products").
				Field("created_by").
				Unique(),

			edge.From("updated_by_user", User.Type).
				Ref("updated_products").
				Field("updated_by").
				Unique(),

			edge.From("deleted_by_user", User.Type).
				Ref("deleted_products").
				Field("deleted_by").
				Unique(),
		}
	default:
		return nil
	}
}

// Hooks of the BaseMixin.
// func (d BaseMixin) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		hook.On(
// 			func(next ent.Mutator) ent.Mutator {
// 				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
// 					// Skip soft-delete, means delete the entity permanently.
// 					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
// 						return next.Mutate(ctx, m)
// 					}
// 					mx, ok := m.(interface {
// 						SetOp(ent.Op)
// 						Client() *gen.Client
// 						SetDeletedAt(time.Time)
// 						WhereP(...func(*sql.Selector))
// 					})
// 					if !ok {
// 						return nil, fmt.Errorf("unexpected mutation type %T", m)
// 					}
// 					d.P(mx)
// 					mx.SetOp(ent.OpUpdate)
// 					mx.SetDeletedAt(time.Now())
// 					return mx.Client().Mutate(ctx, m)
// 				})
// 			},
// 			ent.OpDeleteOne|ent.OpDelete,
// 		),
// 	}
// }

// P adds a storage-level predicate to the queries and mutations.
func (d BaseMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[2].Descriptor().Name), // Assume it's deleted_at
	)
}
