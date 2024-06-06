package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Token struct {
	TokenID string `json:"token_id"`
	Value   string `json:"value"`
}

type CustomerTokensDTO struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Address         string     `json:"address"`
	MinterAddress   string     `json:"minter_address"`
	YielderAddress  string     `json:"yielder_address"`
	OffseterAddress string     `json:"offseter_address"`
	Abi             ProjectAbi `json:"abi"`
	Image           string     `json:"image"`
	Tokens          []Token    `json:"tokens"`
	Slot            int        `json:"slot"`
}

// CustomerTokens holds the schema definition for the CustomerTokens entity.
type CustomerTokens struct {
	ent.Schema
}

// Fields of the CustomerTokens.
func (CustomerTokens) Fields() []ent.Field {
	return []ent.Field{
		field.String("address"),
		field.String("project_address"),
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
