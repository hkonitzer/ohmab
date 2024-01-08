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
	"github.com/hkonitzer/ohmab/ent/privacy"
	"github.com/hkonitzer/ohmab/ent/schema/constants"
	"github.com/hkonitzer/ohmab/ent/schema/hooks"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy/rule"
)

// Business holds the schema definition for the Business entity.
type Business struct {
	ent.Schema
	hooks.AuditLog
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
		field.Text("alias").
			Unique().
			MaxLen(20).
			Annotations(entsql.Annotation{ // If using a SQL-database: change the underlying data type to varchar(20).
				Size: 20,
			},
				entgql.OrderField("ALIAS")).
			Comment("The unqiue alias of the business (short name)"),
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
		edge.To("users", User.Type),
		edge.To("operators", Operator.Type),
	}
}

func (Business) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
	}
}

func (Business) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name1", "name2"),
		index.Fields("alias"),
		index.Fields("email"),
	}
}

func (b Business) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.UpperCaseForBussinessFields(),
		// @TODO: Hook to create public users while adding users to businesses: Update().AddUsers(u)
		b.AuditLogForBusiness(),
	}
}

// Policy defines the privacy policy of the User.
func (Business) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
