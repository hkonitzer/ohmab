package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimetableTypeMixin struct {
	mixin.Schema
}

func (TimetableTypeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("timetable_type").
			Values("DEFAULT", "REGULAR", "CLOSED", "EMERGENCYSERVICE", "HOLIDAY", "SPECIAL").
			Default("DEFAULT").Comment("The type of the timetable entry"),
	}
}
