package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/privacy"
	"github.com/hkonitzer/ohmab/ent/schema/constants"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy/rule"
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
		field.JSON("entity_values", map[string]any{}).
			Immutable().Optional(),
		field.String("entity_uuid").
			Immutable().Optional(),
		field.Time("timestamp").Annotations(entgql.OrderField("timestamp")).
			Default(time.Now).Immutable(),
	}
}

func (AuditLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
	}
}

func (AuditLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("entity_uuid"),
		index.Fields("timestamp"),
	}
}
func (AuditLog) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.DenyIfNoViewer(),
		},
	}
}
