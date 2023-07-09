package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/schema/constants"
	"hynie.de/ohmab/ent/schema/hooks"
)

// Content holds the schema definition for the Content entity.
type Content struct {
	ent.Schema
	hooks.AuditLog
}

// Fields of the Content.
func (Content) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.Enum("type").
			Values("TEXT", "HTML", "CSS", "OTHER").Default("TEXT").
			Comment("The type of the content, only HTML is supported at the moment"),
		field.Text("locale").
			MinLen(5).MaxLen(5).
			Match(localeRex).
			Default("en_US").Comment("The ICU locale identifier for this content, e.g. en_US, de_DE, ..."),
		field.Enum("location").
			Values("TOP", "BOTTOM").Default("TOP"),
		field.Text("content"),
		field.Enum("status").
			Values("DRAFT", "PUBLISHED").Default("DRAFT"),
		field.Time("published_at").
			Annotations(entgql.OrderField("published_at")).
			Optional(),
	}
}

// Edges of the Content.
func (Content) Edges() []ent.Edge {
	return nil
}

// Mixin Time mixin for the Content
func (Content) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		TimetableTypeMixin{},
	}
}

func (Content) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status", "type", "location", "locale").Unique(), // only one content allowed
	}
}

func (c Content) Hooks() []ent.Hook {
	return []ent.Hook{
		c.AuditLogForContent(),
	}
}
