package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/schema/constants"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of an Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New).Immutable(),
		field.Text("addition").
			Optional().Comment("The address addition"),
		field.Text("street").
			Optional(),
		field.Text("city").
			Optional(),
		field.Text("zip").
			Optional(),
		field.Text("state").
			Optional(),
		field.Text("country").
			Optional(),
		field.Text("telephone").
			Optional().Unique().Comment("Telephone number"),
		field.Text("comment").
			Optional().Comment("A comment for this address"),
	}
}

// Mixin Time mixin for the Addresses
func (Address) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).
			Ref("addresses"),
		edge.To("timetables", Timetable.Type),
	}
}
