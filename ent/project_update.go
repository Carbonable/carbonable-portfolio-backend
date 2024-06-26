// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carbonable/carbonable-portfolio-backend/ent/customertokens"
	"github.com/carbonable/carbonable-portfolio-backend/ent/predicate"
	"github.com/carbonable/carbonable-portfolio-backend/ent/project"
	"github.com/carbonable/carbonable-portfolio-backend/internal/model"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAddress sets the "address" field.
func (pu *ProjectUpdate) SetAddress(s string) *ProjectUpdate {
	pu.mutation.SetAddress(s)
	return pu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableAddress(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetAddress(*s)
	}
	return pu
}

// SetSlot sets the "slot" field.
func (pu *ProjectUpdate) SetSlot(i int) *ProjectUpdate {
	pu.mutation.ResetSlot()
	pu.mutation.SetSlot(i)
	return pu
}

// SetNillableSlot sets the "slot" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableSlot(i *int) *ProjectUpdate {
	if i != nil {
		pu.SetSlot(*i)
	}
	return pu
}

// AddSlot adds i to the "slot" field.
func (pu *ProjectUpdate) AddSlot(i int) *ProjectUpdate {
	pu.mutation.AddSlot(i)
	return pu
}

// SetMinterAddress sets the "minter_address" field.
func (pu *ProjectUpdate) SetMinterAddress(s string) *ProjectUpdate {
	pu.mutation.SetMinterAddress(s)
	return pu
}

// SetNillableMinterAddress sets the "minter_address" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableMinterAddress(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetMinterAddress(*s)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableName(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetAbi sets the "abi" field.
func (pu *ProjectUpdate) SetAbi(ma model.ProjectAbi) *ProjectUpdate {
	pu.mutation.SetAbi(ma)
	return pu
}

// SetNillableAbi sets the "abi" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableAbi(ma *model.ProjectAbi) *ProjectUpdate {
	if ma != nil {
		pu.SetAbi(*ma)
	}
	return pu
}

// SetSlotURI sets the "slot_uri" field.
func (pu *ProjectUpdate) SetSlotURI(mu model.SlotUri) *ProjectUpdate {
	pu.mutation.SetSlotURI(mu)
	return pu
}

// SetNillableSlotURI sets the "slot_uri" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableSlotURI(mu *model.SlotUri) *ProjectUpdate {
	if mu != nil {
		pu.SetSlotURI(*mu)
	}
	return pu
}

// SetImage sets the "image" field.
func (pu *ProjectUpdate) SetImage(s string) *ProjectUpdate {
	pu.mutation.SetImage(s)
	return pu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableImage(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetImage(*s)
	}
	return pu
}

// SetYielderAddress sets the "yielder_address" field.
func (pu *ProjectUpdate) SetYielderAddress(s string) *ProjectUpdate {
	pu.mutation.SetYielderAddress(s)
	return pu
}

// SetNillableYielderAddress sets the "yielder_address" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableYielderAddress(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetYielderAddress(*s)
	}
	return pu
}

// ClearYielderAddress clears the value of the "yielder_address" field.
func (pu *ProjectUpdate) ClearYielderAddress() *ProjectUpdate {
	pu.mutation.ClearYielderAddress()
	return pu
}

// SetOffseterAddress sets the "offseter_address" field.
func (pu *ProjectUpdate) SetOffseterAddress(s string) *ProjectUpdate {
	pu.mutation.SetOffseterAddress(s)
	return pu
}

// SetNillableOffseterAddress sets the "offseter_address" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableOffseterAddress(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetOffseterAddress(*s)
	}
	return pu
}

// ClearOffseterAddress clears the value of the "offseter_address" field.
func (pu *ProjectUpdate) ClearOffseterAddress() *ProjectUpdate {
	pu.mutation.ClearOffseterAddress()
	return pu
}

// AddTokenIDs adds the "tokens" edge to the CustomerTokens entity by IDs.
func (pu *ProjectUpdate) AddTokenIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddTokenIDs(ids...)
	return pu
}

// AddTokens adds the "tokens" edges to the CustomerTokens entity.
func (pu *ProjectUpdate) AddTokens(c ...*CustomerTokens) *ProjectUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddTokenIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearTokens clears all "tokens" edges to the CustomerTokens entity.
func (pu *ProjectUpdate) ClearTokens() *ProjectUpdate {
	pu.mutation.ClearTokens()
	return pu
}

// RemoveTokenIDs removes the "tokens" edge to CustomerTokens entities by IDs.
func (pu *ProjectUpdate) RemoveTokenIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveTokenIDs(ids...)
	return pu
}

// RemoveTokens removes "tokens" edges to CustomerTokens entities.
func (pu *ProjectUpdate) RemoveTokens(c ...*CustomerTokens) *ProjectUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveTokenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Address(); ok {
		_spec.SetField(project.FieldAddress, field.TypeString, value)
	}
	if value, ok := pu.mutation.Slot(); ok {
		_spec.SetField(project.FieldSlot, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedSlot(); ok {
		_spec.AddField(project.FieldSlot, field.TypeInt, value)
	}
	if value, ok := pu.mutation.MinterAddress(); ok {
		_spec.SetField(project.FieldMinterAddress, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Abi(); ok {
		_spec.SetField(project.FieldAbi, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.SlotURI(); ok {
		_spec.SetField(project.FieldSlotURI, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.Image(); ok {
		_spec.SetField(project.FieldImage, field.TypeString, value)
	}
	if value, ok := pu.mutation.YielderAddress(); ok {
		_spec.SetField(project.FieldYielderAddress, field.TypeString, value)
	}
	if pu.mutation.YielderAddressCleared() {
		_spec.ClearField(project.FieldYielderAddress, field.TypeString)
	}
	if value, ok := pu.mutation.OffseterAddress(); ok {
		_spec.SetField(project.FieldOffseterAddress, field.TypeString, value)
	}
	if pu.mutation.OffseterAddressCleared() {
		_spec.ClearField(project.FieldOffseterAddress, field.TypeString)
	}
	if pu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.TokensTable,
			Columns: project.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customertokens.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTokensIDs(); len(nodes) > 0 && !pu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.TokensTable,
			Columns: project.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customertokens.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.TokensTable,
			Columns: project.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customertokens.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetAddress sets the "address" field.
func (puo *ProjectUpdateOne) SetAddress(s string) *ProjectUpdateOne {
	puo.mutation.SetAddress(s)
	return puo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableAddress(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetAddress(*s)
	}
	return puo
}

// SetSlot sets the "slot" field.
func (puo *ProjectUpdateOne) SetSlot(i int) *ProjectUpdateOne {
	puo.mutation.ResetSlot()
	puo.mutation.SetSlot(i)
	return puo
}

// SetNillableSlot sets the "slot" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableSlot(i *int) *ProjectUpdateOne {
	if i != nil {
		puo.SetSlot(*i)
	}
	return puo
}

// AddSlot adds i to the "slot" field.
func (puo *ProjectUpdateOne) AddSlot(i int) *ProjectUpdateOne {
	puo.mutation.AddSlot(i)
	return puo
}

// SetMinterAddress sets the "minter_address" field.
func (puo *ProjectUpdateOne) SetMinterAddress(s string) *ProjectUpdateOne {
	puo.mutation.SetMinterAddress(s)
	return puo
}

// SetNillableMinterAddress sets the "minter_address" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableMinterAddress(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetMinterAddress(*s)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableName(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetAbi sets the "abi" field.
func (puo *ProjectUpdateOne) SetAbi(ma model.ProjectAbi) *ProjectUpdateOne {
	puo.mutation.SetAbi(ma)
	return puo
}

// SetNillableAbi sets the "abi" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableAbi(ma *model.ProjectAbi) *ProjectUpdateOne {
	if ma != nil {
		puo.SetAbi(*ma)
	}
	return puo
}

// SetSlotURI sets the "slot_uri" field.
func (puo *ProjectUpdateOne) SetSlotURI(mu model.SlotUri) *ProjectUpdateOne {
	puo.mutation.SetSlotURI(mu)
	return puo
}

// SetNillableSlotURI sets the "slot_uri" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableSlotURI(mu *model.SlotUri) *ProjectUpdateOne {
	if mu != nil {
		puo.SetSlotURI(*mu)
	}
	return puo
}

// SetImage sets the "image" field.
func (puo *ProjectUpdateOne) SetImage(s string) *ProjectUpdateOne {
	puo.mutation.SetImage(s)
	return puo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableImage(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetImage(*s)
	}
	return puo
}

// SetYielderAddress sets the "yielder_address" field.
func (puo *ProjectUpdateOne) SetYielderAddress(s string) *ProjectUpdateOne {
	puo.mutation.SetYielderAddress(s)
	return puo
}

// SetNillableYielderAddress sets the "yielder_address" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableYielderAddress(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetYielderAddress(*s)
	}
	return puo
}

// ClearYielderAddress clears the value of the "yielder_address" field.
func (puo *ProjectUpdateOne) ClearYielderAddress() *ProjectUpdateOne {
	puo.mutation.ClearYielderAddress()
	return puo
}

// SetOffseterAddress sets the "offseter_address" field.
func (puo *ProjectUpdateOne) SetOffseterAddress(s string) *ProjectUpdateOne {
	puo.mutation.SetOffseterAddress(s)
	return puo
}

// SetNillableOffseterAddress sets the "offseter_address" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableOffseterAddress(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetOffseterAddress(*s)
	}
	return puo
}

// ClearOffseterAddress clears the value of the "offseter_address" field.
func (puo *ProjectUpdateOne) ClearOffseterAddress() *ProjectUpdateOne {
	puo.mutation.ClearOffseterAddress()
	return puo
}

// AddTokenIDs adds the "tokens" edge to the CustomerTokens entity by IDs.
func (puo *ProjectUpdateOne) AddTokenIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddTokenIDs(ids...)
	return puo
}

// AddTokens adds the "tokens" edges to the CustomerTokens entity.
func (puo *ProjectUpdateOne) AddTokens(c ...*CustomerTokens) *ProjectUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddTokenIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearTokens clears all "tokens" edges to the CustomerTokens entity.
func (puo *ProjectUpdateOne) ClearTokens() *ProjectUpdateOne {
	puo.mutation.ClearTokens()
	return puo
}

// RemoveTokenIDs removes the "tokens" edge to CustomerTokens entities by IDs.
func (puo *ProjectUpdateOne) RemoveTokenIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveTokenIDs(ids...)
	return puo
}

// RemoveTokens removes "tokens" edges to CustomerTokens entities.
func (puo *ProjectUpdateOne) RemoveTokens(c ...*CustomerTokens) *ProjectUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveTokenIDs(ids...)
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Address(); ok {
		_spec.SetField(project.FieldAddress, field.TypeString, value)
	}
	if value, ok := puo.mutation.Slot(); ok {
		_spec.SetField(project.FieldSlot, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedSlot(); ok {
		_spec.AddField(project.FieldSlot, field.TypeInt, value)
	}
	if value, ok := puo.mutation.MinterAddress(); ok {
		_spec.SetField(project.FieldMinterAddress, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Abi(); ok {
		_spec.SetField(project.FieldAbi, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.SlotURI(); ok {
		_spec.SetField(project.FieldSlotURI, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.Image(); ok {
		_spec.SetField(project.FieldImage, field.TypeString, value)
	}
	if value, ok := puo.mutation.YielderAddress(); ok {
		_spec.SetField(project.FieldYielderAddress, field.TypeString, value)
	}
	if puo.mutation.YielderAddressCleared() {
		_spec.ClearField(project.FieldYielderAddress, field.TypeString)
	}
	if value, ok := puo.mutation.OffseterAddress(); ok {
		_spec.SetField(project.FieldOffseterAddress, field.TypeString, value)
	}
	if puo.mutation.OffseterAddressCleared() {
		_spec.ClearField(project.FieldOffseterAddress, field.TypeString)
	}
	if puo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.TokensTable,
			Columns: project.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customertokens.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTokensIDs(); len(nodes) > 0 && !puo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.TokensTable,
			Columns: project.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customertokens.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.TokensTable,
			Columns: project.TokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customertokens.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
