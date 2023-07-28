package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/schema/constants"
	"github.com/hkonitzer/ohmab/ent/schema/hooks"
)

// Timetable holds the schema definition for a Timetable entity.
type Timetable struct {
	ent.Schema
	hooks.AuditLog
}

// Fields of a Timetable.
func (Timetable) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.Time("datetime_from").
			Annotations(entgql.OrderField("datetimeFrom")),
		field.Int("duration").
			Positive().Range(1, 24).Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "TINYINT",  // Override MySQL.
				dialect.Postgres: "SMALLINT", // Override Postgres.
				dialect.SQLite:   "INTEGER",  // Override SQLite.
			}).Comment("The duration of the timetable entry in hours, overwrites datetime_to"),
		field.Time("datetime_to").
			Optional().Comment("The end of the timetable entry, only used if duration is not set"),
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
		TimetableTypeMixin{},
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
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
	}
}

func (Timetable) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("datetime_from"),
	}
}
func (t Timetable) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.EnsureDurationIsSet(),
		hooks.EnsureOnlyOneTimetableEntry(),
		t.AuditLogForTimetable(),
	}
}
