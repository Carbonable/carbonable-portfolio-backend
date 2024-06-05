// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/carbonable/carbonable-portfolio-backend/ent/project"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Slot holds the value of the "slot" field.
	Slot int `json:"slot,omitempty"`
	// MinterAddress holds the value of the "minter_address" field.
	MinterAddress string `json:"minter_address,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Abi holds the value of the "abi" field.
	Abi json.RawMessage `json:"abi,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// YielderAddress holds the value of the "yielder_address" field.
	YielderAddress string `json:"yielder_address,omitempty"`
	// OffseterAddress holds the value of the "offseter_address" field.
	OffseterAddress string `json:"offseter_address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges        ProjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Tokens holds the value of the tokens edge.
	Tokens []*CustomerTokens `json:"tokens,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedTokens map[string][]*CustomerTokens
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) TokensOrErr() ([]*CustomerTokens, error) {
	if e.loadedTypes[0] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldAbi:
			values[i] = new([]byte)
		case project.FieldID, project.FieldSlot:
			values[i] = new(sql.NullInt64)
		case project.FieldAddress, project.FieldMinterAddress, project.FieldName, project.FieldImage, project.FieldYielderAddress, project.FieldOffseterAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case project.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				pr.Address = value.String
			}
		case project.FieldSlot:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field slot", values[i])
			} else if value.Valid {
				pr.Slot = int(value.Int64)
			}
		case project.FieldMinterAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field minter_address", values[i])
			} else if value.Valid {
				pr.MinterAddress = value.String
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldAbi:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field abi", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Abi); err != nil {
					return fmt.Errorf("unmarshal field abi: %w", err)
				}
			}
		case project.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				pr.Image = value.String
			}
		case project.FieldYielderAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field yielder_address", values[i])
			} else if value.Valid {
				pr.YielderAddress = value.String
			}
		case project.FieldOffseterAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field offseter_address", values[i])
			} else if value.Valid {
				pr.OffseterAddress = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Project.
// This includes values selected through modifiers, order, etc.
func (pr *Project) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryTokens queries the "tokens" edge of the Project entity.
func (pr *Project) QueryTokens() *CustomerTokensQuery {
	return NewProjectClient(pr.config).QueryTokens(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("address=")
	builder.WriteString(pr.Address)
	builder.WriteString(", ")
	builder.WriteString("slot=")
	builder.WriteString(fmt.Sprintf("%v", pr.Slot))
	builder.WriteString(", ")
	builder.WriteString("minter_address=")
	builder.WriteString(pr.MinterAddress)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("abi=")
	builder.WriteString(fmt.Sprintf("%v", pr.Abi))
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(pr.Image)
	builder.WriteString(", ")
	builder.WriteString("yielder_address=")
	builder.WriteString(pr.YielderAddress)
	builder.WriteString(", ")
	builder.WriteString("offseter_address=")
	builder.WriteString(pr.OffseterAddress)
	builder.WriteByte(')')
	return builder.String()
}

// NamedTokens returns the Tokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Project) NamedTokens(name string) ([]*CustomerTokens, error) {
	if pr.Edges.namedTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Project) appendNamedTokens(name string, edges ...*CustomerTokens) {
	if pr.Edges.namedTokens == nil {
		pr.Edges.namedTokens = make(map[string][]*CustomerTokens)
	}
	if len(edges) == 0 {
		pr.Edges.namedTokens[name] = []*CustomerTokens{}
	} else {
		pr.Edges.namedTokens[name] = append(pr.Edges.namedTokens[name], edges...)
	}
}

// Projects is a parsable slice of Project.
type Projects []*Project
