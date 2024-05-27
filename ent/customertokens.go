// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/carbonable/carbonable-portfolio-backend/ent/customertokens"
)

// CustomerTokens is the model entity for the CustomerTokens schema.
type CustomerTokens struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Slot holds the value of the "slot" field.
	Slot int `json:"slot,omitempty"`
	// TokenID holds the value of the "token_id" field.
	TokenID string `json:"token_id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerTokensQuery when eager-loading is set.
	Edges        CustomerTokensEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CustomerTokensEdges holds the relations/edges for other nodes in the graph.
type CustomerTokensEdges struct {
	// Project holds the value of the project edge.
	Project []*Project `json:"project,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerTokensEdges) ProjectOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CustomerTokens) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customertokens.FieldID, customertokens.FieldSlot:
			values[i] = new(sql.NullInt64)
		case customertokens.FieldAddress, customertokens.FieldTokenID, customertokens.FieldValue:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CustomerTokens fields.
func (ct *CustomerTokens) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customertokens.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ct.ID = int(value.Int64)
		case customertokens.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				ct.Address = value.String
			}
		case customertokens.FieldSlot:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field slot", values[i])
			} else if value.Valid {
				ct.Slot = int(value.Int64)
			}
		case customertokens.FieldTokenID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_id", values[i])
			} else if value.Valid {
				ct.TokenID = value.String
			}
		case customertokens.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ct.Value = value.String
			}
		default:
			ct.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the CustomerTokens.
// This includes values selected through modifiers, order, etc.
func (ct *CustomerTokens) GetValue(name string) (ent.Value, error) {
	return ct.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the CustomerTokens entity.
func (ct *CustomerTokens) QueryProject() *ProjectQuery {
	return NewCustomerTokensClient(ct.config).QueryProject(ct)
}

// Update returns a builder for updating this CustomerTokens.
// Note that you need to call CustomerTokens.Unwrap() before calling this method if this CustomerTokens
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *CustomerTokens) Update() *CustomerTokensUpdateOne {
	return NewCustomerTokensClient(ct.config).UpdateOne(ct)
}

// Unwrap unwraps the CustomerTokens entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *CustomerTokens) Unwrap() *CustomerTokens {
	_tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("ent: CustomerTokens is not a transactional entity")
	}
	ct.config.driver = _tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *CustomerTokens) String() string {
	var builder strings.Builder
	builder.WriteString("CustomerTokens(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ct.ID))
	builder.WriteString("address=")
	builder.WriteString(ct.Address)
	builder.WriteString(", ")
	builder.WriteString("slot=")
	builder.WriteString(fmt.Sprintf("%v", ct.Slot))
	builder.WriteString(", ")
	builder.WriteString("token_id=")
	builder.WriteString(ct.TokenID)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(ct.Value)
	builder.WriteByte(')')
	return builder.String()
}

// CustomerTokensSlice is a parsable slice of CustomerTokens.
type CustomerTokensSlice []*CustomerTokens
