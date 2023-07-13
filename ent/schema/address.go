package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/schema/constants"
	"github.com/hkonitzer/ohmab/ent/schema/hooks"
	"regexp"
)

var localeRex, _ = regexp.Compile("^[a-z]{2,4}(_[A-Z][a-z]{3})?(_([A-Z]{2}|[0-9]{3}))?$")

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
	hooks.AuditLog
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
		field.Text("locale").
			MinLen(5).MaxLen(5).
			Match(localeRex).
			Default("en_US").Comment("The ICU locale identifier of the address, e.g. en_US, de_DE, ..."),
		field.Bool("primary").
			Default(false).Comment("Is this the primary address?"),
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

func (Address) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).
			Ref("addresses").Unique(),
		edge.To("timetables", Timetable.Type),
	}
}

func (a Address) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.EnsureOnlyOnePrimaryAddress(),
		a.AuditLogForAddress(),
	}
}
