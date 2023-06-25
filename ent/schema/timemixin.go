package schema

import (
	"hynie.de/ohmab/ent/schema/constants"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// -------------------------------------------------
// Mixin definition

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time(constants.TimeMixinCreatedAtFieldName).
			Immutable().
			Default(time.Now),
		field.Time(constants.TimeMixinUpdatedAtFieldName).
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time(constants.TimeMixinDeletedAtFieldName).
			Optional(),
	}
}
