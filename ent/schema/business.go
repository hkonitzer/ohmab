package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/schema/constants"
	"hynie.de/ohmab/ent/schema/hooks"
)

// Business holds the schema definition for the Business entity.
type Business struct {
	ent.Schema
}

// Fields of the Business.
func (Business) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(constants.IDFieldName, uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.Text("name1").
			NotEmpty().Annotations(entgql.OrderField("NAME1")).Comment("The main name of the business"),
		field.Text("name2").
			Optional().Comment("The optional second name of the business"),
		field.Text("telephone").
			Optional().Unique().Comment("Telephone number"),
		field.Text("email").
			Optional().Unique().Comment("Email address (has to be unique)"),
		field.Text("website").
			Optional().Comment("Website address"),
		field.Text("comment").
			Optional().Comment("A comment for this business"),
		field.Bool(constants.ActiveFieldName).Default(true).Comment("Is the business active?"),
	}
}

// Mixin Time for the business
func (Business) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Business.
func (Business) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("addresses", Address.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("tags", Tag.Type),
		edge.From("users", User.Type).Ref("businesses").Unique(),
	}
}

func (Business) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
	}
}

func (Business) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name1", "name2"),
		index.Fields("email"),
	}
}

func (Business) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.AuditLogForBusiness(),
	}
}
