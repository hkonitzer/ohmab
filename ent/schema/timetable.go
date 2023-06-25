package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/schema/constants"
	"hynie.de/ohmab/ent/schema/hooks"
)

// Timetable holds the schema definition for a Timetable entity.
type Timetable struct {
	ent.Schema
}

// Fields of a Timetable.
func (Timetable) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.Enum("type").
			Values("DEFAULT", "REGULAR", "CLOSED", "EMERGENCYSERVICE", "HOLIDAY", "SPECIAL").Default("DEFAULT"),
		field.Time("datetime_from").Optional(),
		field.Time("datetime_to").Optional(),
		field.Bool("time_whole_day").Default(false),
		field.Text("comment").Optional(),
		field.Text("availability_by_phone").Optional(),
		field.Text("availability_by_email").Optional(),
		field.Text("availability_by_sms").Optional(),
		field.Text("availability_by_whatsapp").Optional(),
	}
}

// Mixin Time mixin for a Timetable
func (Timetable) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of a Timetable.
func (Timetable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("address", Address.Type).
			Required().
			Unique().
			Ref("timetables"),
		edge.To("users_on_duty", User.Type),
	}
}

func (Timetable) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
	}
}

func (Timetable) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("datetime_from"),
	}
}
func (Timetable) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.AuditLogForTimetable(),
	}
}
