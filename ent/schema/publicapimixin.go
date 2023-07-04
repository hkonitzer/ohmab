package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type PublicApiMixin struct {
	mixin.Schema
}

const UsePublicApiValue = "1"

func (PublicApiMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Text("use_publicapi"). // bool fields cannot be sensitive
						Default("").Sensitive().
						Comment("Set to 1 if this record should be visible in the public api"), // if empty = no
	}
}

func (PublicApiMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("use_publicapi"),
	}
}
