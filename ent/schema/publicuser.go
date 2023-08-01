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

// User holds the schema definition for the User entity.
type PublicUser struct {
	ent.Schema
}

// Fields of a User.
func (PublicUser) Fields() []ent.Field {
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

func (PublicUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("businesses", Business.Type).Ref("public_users").Comment("The businesses this user is associated with"),
		edge.From("timetable", Timetable.Type).Ref("users_on_duty").Comment("The persons on duty for this timetable entry"),
	}
}

func (PublicUser) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{ // see also user.go
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
