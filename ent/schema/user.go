package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/schema/constants"
	"hynie.de/ohmab/ent/schema/hooks"
	"hynie.de/ohmab/internal/pkg/privacy"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of a User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.Text("login").
			Sensitive().NotEmpty(),
		field.Text("surname").
			NotEmpty().Annotations(entgql.OrderField("SURNAME")).Comment("The surname of a user"),
		field.Text("firstname").
			Comment("The first name of a user"),
		field.Text("title").
			Optional().Comment("The title of a user like PhD"),
		field.Text("email").
			Unique(),
		field.Text("passwordhash").
			Optional().Sensitive().Comment("The password hash of a user"),
		field.Text("comment").
			Optional().Comment("A comment for this user"),
		field.Bool(constants.ActiveFieldName).Default(true).Comment("Is the user active?"),
		field.String("role").
			Sensitive().Default(privacy.ViewerRoleAsString()).Comment("The role of the user"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		PublicApiMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("businesses", Business.Type).Ref("users").Comment("The businesses this user is associated with"),
		edge.To("tags", Tag.Type),
		edge.From("timetable", Timetable.Type).Ref("users_on_duty").Comment("The persons on duty for this timetable entry"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.AuditLogForUser(),
	}
}
