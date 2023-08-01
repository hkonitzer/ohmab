package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/privacy"
	"github.com/hkonitzer/ohmab/ent/schema/constants"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy/rule"
)

// Operator holds the schema definition for the Operator entity.
// These users were derived from the users if the flag public_api is set to 1
// There is no connection to the Users table, because of the public api
// See hooks.UpdatePublicUser()
type Operator struct {
	ent.Schema
}

// Fields of a User.
func (Operator) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.Text("surname").
			NotEmpty().Annotations(entgql.OrderField("SURNAME")).Comment("The surname of a user"),
		field.Text("firstname").
			Comment("The first name of a user"),
		field.Text("title").
			Optional().Comment("The title of a user like PhD"),
		field.Text("email").
			Sensitive().Unique(),
	}
}

func (Operator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("businesses", Business.Type).Ref("operators").Comment("The businesses this user is associated with"),
		edge.From("timetable", Timetable.Type).Ref("operators_on_duty").Comment("The persons on duty for this timetable entry"),
	}
}

func (Operator) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{ // see also user.go
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
