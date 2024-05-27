package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CustomerTokens holds the schema definition for the CustomerTokens entity.
type CustomerTokens struct {
	ent.Schema
}

// Fields of the CustomerTokens.
func (CustomerTokens) Fields() []ent.Field {
	return []ent.Field{
		field.String("address"),
		field.Int("slot"),
		field.String("token_id"),
		field.String("value"),
	}
}

// Edges of the CustomerTokens.
func (CustomerTokens) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("tokens"),
	}
}

func (CustomerTokens) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token_id", "slot").Unique(),
	}
}
