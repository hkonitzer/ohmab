package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/schema/constants"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.Text("name").
			NotEmpty().Annotations(entgql.OrderField("NAME")).Comment("The  name of the tag"),
		field.Text("comment").
			Optional().Comment("A comment for this tag"),
	}
}

// Mixin Time mixin for the Tags
func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of a Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).Ref("tags"),
		edge.From("user", User.Type).Ref("tags"),
	}
}
