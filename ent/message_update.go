// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/velcro-xiv/velcro/ent/message"
	"github.com/velcro-xiv/velcro/ent/predicate"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetTimestamp sets the "timestamp" field.
func (mu *MessageUpdate) SetTimestamp(t time.Time) *MessageUpdate {
	mu.mutation.SetTimestamp(t)
	return mu
}

// SetVersion sets the "version" field.
func (mu *MessageUpdate) SetVersion(i int) *MessageUpdate {
	mu.mutation.ResetVersion()
	mu.mutation.SetVersion(i)
	return mu
}

// AddVersion adds i to the "version" field.
func (mu *MessageUpdate) AddVersion(i int) *MessageUpdate {
	mu.mutation.AddVersion(i)
	return mu
}

// SetSegment sets the "segment" field.
func (mu *MessageUpdate) SetSegment(i int) *MessageUpdate {
	mu.mutation.ResetSegment()
	mu.mutation.SetSegment(i)
	return mu
}

// SetNillableSegment sets the "segment" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableSegment(i *int) *MessageUpdate {
	if i != nil {
		mu.SetSegment(*i)
	}
	return mu
}

// AddSegment adds i to the "segment" field.
func (mu *MessageUpdate) AddSegment(i int) *MessageUpdate {
	mu.mutation.AddSegment(i)
	return mu
}

// SetOpcode sets the "opcode" field.
func (mu *MessageUpdate) SetOpcode(i int) *MessageUpdate {
	mu.mutation.ResetOpcode()
	mu.mutation.SetOpcode(i)
	return mu
}

// SetNillableOpcode sets the "opcode" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableOpcode(i *int) *MessageUpdate {
	if i != nil {
		mu.SetOpcode(*i)
	}
	return mu
}

// AddOpcode adds i to the "opcode" field.
func (mu *MessageUpdate) AddOpcode(i int) *MessageUpdate {
	mu.mutation.AddOpcode(i)
	return mu
}

// ClearOpcode clears the value of the "opcode" field.
func (mu *MessageUpdate) ClearOpcode() *MessageUpdate {
	mu.mutation.ClearOpcode()
	return mu
}

// SetSourceAddress sets the "source_address" field.
func (mu *MessageUpdate) SetSourceAddress(s string) *MessageUpdate {
	mu.mutation.SetSourceAddress(s)
	return mu
}

// SetSourcePort sets the "source_port" field.
func (mu *MessageUpdate) SetSourcePort(i int) *MessageUpdate {
	mu.mutation.ResetSourcePort()
	mu.mutation.SetSourcePort(i)
	return mu
}

// AddSourcePort adds i to the "source_port" field.
func (mu *MessageUpdate) AddSourcePort(i int) *MessageUpdate {
	mu.mutation.AddSourcePort(i)
	return mu
}

// SetDestinationAddress sets the "destination_address" field.
func (mu *MessageUpdate) SetDestinationAddress(s string) *MessageUpdate {
	mu.mutation.SetDestinationAddress(s)
	return mu
}

// SetDestinationPort sets the "destination_port" field.
func (mu *MessageUpdate) SetDestinationPort(i int) *MessageUpdate {
	mu.mutation.ResetDestinationPort()
	mu.mutation.SetDestinationPort(i)
	return mu
}

// AddDestinationPort adds i to the "destination_port" field.
func (mu *MessageUpdate) AddDestinationPort(i int) *MessageUpdate {
	mu.mutation.AddDestinationPort(i)
	return mu
}

// SetData sets the "data" field.
func (mu *MessageUpdate) SetData(b []byte) *MessageUpdate {
	mu.mutation.SetData(b)
	return mu
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldTimestamp,
		})
	}
	if value, ok := mu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldVersion,
		})
	}
	if value, ok := mu.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldVersion,
		})
	}
	if value, ok := mu.mutation.Segment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSegment,
		})
	}
	if value, ok := mu.mutation.AddedSegment(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSegment,
		})
	}
	if value, ok := mu.mutation.Opcode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldOpcode,
		})
	}
	if value, ok := mu.mutation.AddedOpcode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldOpcode,
		})
	}
	if mu.mutation.OpcodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: message.FieldOpcode,
		})
	}
	if value, ok := mu.mutation.SourceAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldSourceAddress,
		})
	}
	if value, ok := mu.mutation.SourcePort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSourcePort,
		})
	}
	if value, ok := mu.mutation.AddedSourcePort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSourcePort,
		})
	}
	if value, ok := mu.mutation.DestinationAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldDestinationAddress,
		})
	}
	if value, ok := mu.mutation.DestinationPort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldDestinationPort,
		})
	}
	if value, ok := mu.mutation.AddedDestinationPort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldDestinationPort,
		})
	}
	if value, ok := mu.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: message.FieldData,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetTimestamp sets the "timestamp" field.
func (muo *MessageUpdateOne) SetTimestamp(t time.Time) *MessageUpdateOne {
	muo.mutation.SetTimestamp(t)
	return muo
}

// SetVersion sets the "version" field.
func (muo *MessageUpdateOne) SetVersion(i int) *MessageUpdateOne {
	muo.mutation.ResetVersion()
	muo.mutation.SetVersion(i)
	return muo
}

// AddVersion adds i to the "version" field.
func (muo *MessageUpdateOne) AddVersion(i int) *MessageUpdateOne {
	muo.mutation.AddVersion(i)
	return muo
}

// SetSegment sets the "segment" field.
func (muo *MessageUpdateOne) SetSegment(i int) *MessageUpdateOne {
	muo.mutation.ResetSegment()
	muo.mutation.SetSegment(i)
	return muo
}

// SetNillableSegment sets the "segment" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableSegment(i *int) *MessageUpdateOne {
	if i != nil {
		muo.SetSegment(*i)
	}
	return muo
}

// AddSegment adds i to the "segment" field.
func (muo *MessageUpdateOne) AddSegment(i int) *MessageUpdateOne {
	muo.mutation.AddSegment(i)
	return muo
}

// SetOpcode sets the "opcode" field.
func (muo *MessageUpdateOne) SetOpcode(i int) *MessageUpdateOne {
	muo.mutation.ResetOpcode()
	muo.mutation.SetOpcode(i)
	return muo
}

// SetNillableOpcode sets the "opcode" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableOpcode(i *int) *MessageUpdateOne {
	if i != nil {
		muo.SetOpcode(*i)
	}
	return muo
}

// AddOpcode adds i to the "opcode" field.
func (muo *MessageUpdateOne) AddOpcode(i int) *MessageUpdateOne {
	muo.mutation.AddOpcode(i)
	return muo
}

// ClearOpcode clears the value of the "opcode" field.
func (muo *MessageUpdateOne) ClearOpcode() *MessageUpdateOne {
	muo.mutation.ClearOpcode()
	return muo
}

// SetSourceAddress sets the "source_address" field.
func (muo *MessageUpdateOne) SetSourceAddress(s string) *MessageUpdateOne {
	muo.mutation.SetSourceAddress(s)
	return muo
}

// SetSourcePort sets the "source_port" field.
func (muo *MessageUpdateOne) SetSourcePort(i int) *MessageUpdateOne {
	muo.mutation.ResetSourcePort()
	muo.mutation.SetSourcePort(i)
	return muo
}

// AddSourcePort adds i to the "source_port" field.
func (muo *MessageUpdateOne) AddSourcePort(i int) *MessageUpdateOne {
	muo.mutation.AddSourcePort(i)
	return muo
}

// SetDestinationAddress sets the "destination_address" field.
func (muo *MessageUpdateOne) SetDestinationAddress(s string) *MessageUpdateOne {
	muo.mutation.SetDestinationAddress(s)
	return muo
}

// SetDestinationPort sets the "destination_port" field.
func (muo *MessageUpdateOne) SetDestinationPort(i int) *MessageUpdateOne {
	muo.mutation.ResetDestinationPort()
	muo.mutation.SetDestinationPort(i)
	return muo
}

// AddDestinationPort adds i to the "destination_port" field.
func (muo *MessageUpdateOne) AddDestinationPort(i int) *MessageUpdateOne {
	muo.mutation.AddDestinationPort(i)
	return muo
}

// SetData sets the "data" field.
func (muo *MessageUpdateOne) SetData(b []byte) *MessageUpdateOne {
	muo.mutation.SetData(b)
	return muo
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Message)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldTimestamp,
		})
	}
	if value, ok := muo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldVersion,
		})
	}
	if value, ok := muo.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldVersion,
		})
	}
	if value, ok := muo.mutation.Segment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSegment,
		})
	}
	if value, ok := muo.mutation.AddedSegment(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSegment,
		})
	}
	if value, ok := muo.mutation.Opcode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldOpcode,
		})
	}
	if value, ok := muo.mutation.AddedOpcode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldOpcode,
		})
	}
	if muo.mutation.OpcodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: message.FieldOpcode,
		})
	}
	if value, ok := muo.mutation.SourceAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldSourceAddress,
		})
	}
	if value, ok := muo.mutation.SourcePort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSourcePort,
		})
	}
	if value, ok := muo.mutation.AddedSourcePort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSourcePort,
		})
	}
	if value, ok := muo.mutation.DestinationAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldDestinationAddress,
		})
	}
	if value, ok := muo.mutation.DestinationPort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldDestinationPort,
		})
	}
	if value, ok := muo.mutation.AddedDestinationPort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldDestinationPort,
		})
	}
	if value, ok := muo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: message.FieldData,
		})
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
