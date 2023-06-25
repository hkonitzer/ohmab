package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/schema/constants"
	"time"
)

type AuditLog struct {
	ent.Schema
}

// Fields of the Business.
func (AuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.Text("user").
			Immutable().NotEmpty(),
		field.Text("action").
			Immutable().NotEmpty(),
		field.Text("entity_schema").
			Immutable().NotEmpty(),
		field.JSON("entity_values", map[string]string{}).
			Immutable().Optional().
			Annotations(entgql.Type("Map")),
		field.String("entity_uuid").
			Immutable().Optional(),
		field.Time("timestamp").
			Default(time.Now).Immutable(),
	}
}

func (AuditLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("entity_uuid"),
	}
}
