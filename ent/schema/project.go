package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/carbonable/carbonable-portfolio-backend/internal/model"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("address"),
		field.Int("slot"),
		field.String("minter_address"),
		field.String("name"),
		field.JSON("abi", model.ProjectAbi{}),
		field.JSON("slot_uri", model.SlotUri{}),
		field.Text("image"),
		field.String("yielder_address").Optional(),
		field.String("offseter_address").Optional(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", CustomerTokens.Type),
	}
}

func (Project) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address", "slot").Unique(),
	}
}
