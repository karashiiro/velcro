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
	"github.com/velcro-xiv/velcro/ent/logevent"
	"github.com/velcro-xiv/velcro/ent/predicate"
)

// LogEventUpdate is the builder for updating LogEvent entities.
type LogEventUpdate struct {
	config
	hooks    []Hook
	mutation *LogEventMutation
}

// Where appends a list predicates to the LogEventUpdate builder.
func (leu *LogEventUpdate) Where(ps ...predicate.LogEvent) *LogEventUpdate {
	leu.mutation.Where(ps...)
	return leu
}

// SetTimestamp sets the "timestamp" field.
func (leu *LogEventUpdate) SetTimestamp(t time.Time) *LogEventUpdate {
	leu.mutation.SetTimestamp(t)
	return leu
}

// SetLevel sets the "level" field.
func (leu *LogEventUpdate) SetLevel(i int) *LogEventUpdate {
	leu.mutation.ResetLevel()
	leu.mutation.SetLevel(i)
	return leu
}

// AddLevel adds i to the "level" field.
func (leu *LogEventUpdate) AddLevel(i int) *LogEventUpdate {
	leu.mutation.AddLevel(i)
	return leu
}

// SetMessage sets the "message" field.
func (leu *LogEventUpdate) SetMessage(s string) *LogEventUpdate {
	leu.mutation.SetMessage(s)
	return leu
}

// Mutation returns the LogEventMutation object of the builder.
func (leu *LogEventUpdate) Mutation() *LogEventMutation {
	return leu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (leu *LogEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(leu.hooks) == 0 {
		affected, err = leu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LogEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			leu.mutation = mutation
			affected, err = leu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(leu.hooks) - 1; i >= 0; i-- {
			if leu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = leu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, leu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (leu *LogEventUpdate) SaveX(ctx context.Context) int {
	affected, err := leu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (leu *LogEventUpdate) Exec(ctx context.Context) error {
	_, err := leu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (leu *LogEventUpdate) ExecX(ctx context.Context) {
	if err := leu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (leu *LogEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   logevent.Table,
			Columns: logevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logevent.FieldID,
			},
		},
	}
	if ps := leu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := leu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: logevent.FieldTimestamp,
		})
	}
	if value, ok := leu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: logevent.FieldLevel,
		})
	}
	if value, ok := leu.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: logevent.FieldLevel,
		})
	}
	if value, ok := leu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logevent.FieldMessage,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, leu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LogEventUpdateOne is the builder for updating a single LogEvent entity.
type LogEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LogEventMutation
}

// SetTimestamp sets the "timestamp" field.
func (leuo *LogEventUpdateOne) SetTimestamp(t time.Time) *LogEventUpdateOne {
	leuo.mutation.SetTimestamp(t)
	return leuo
}

// SetLevel sets the "level" field.
func (leuo *LogEventUpdateOne) SetLevel(i int) *LogEventUpdateOne {
	leuo.mutation.ResetLevel()
	leuo.mutation.SetLevel(i)
	return leuo
}

// AddLevel adds i to the "level" field.
func (leuo *LogEventUpdateOne) AddLevel(i int) *LogEventUpdateOne {
	leuo.mutation.AddLevel(i)
	return leuo
}

// SetMessage sets the "message" field.
func (leuo *LogEventUpdateOne) SetMessage(s string) *LogEventUpdateOne {
	leuo.mutation.SetMessage(s)
	return leuo
}

// Mutation returns the LogEventMutation object of the builder.
func (leuo *LogEventUpdateOne) Mutation() *LogEventMutation {
	return leuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (leuo *LogEventUpdateOne) Select(field string, fields ...string) *LogEventUpdateOne {
	leuo.fields = append([]string{field}, fields...)
	return leuo
}

// Save executes the query and returns the updated LogEvent entity.
func (leuo *LogEventUpdateOne) Save(ctx context.Context) (*LogEvent, error) {
	var (
		err  error
		node *LogEvent
	)
	if len(leuo.hooks) == 0 {
		node, err = leuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LogEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			leuo.mutation = mutation
			node, err = leuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(leuo.hooks) - 1; i >= 0; i-- {
			if leuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = leuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, leuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LogEvent)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LogEventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (leuo *LogEventUpdateOne) SaveX(ctx context.Context) *LogEvent {
	node, err := leuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (leuo *LogEventUpdateOne) Exec(ctx context.Context) error {
	_, err := leuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (leuo *LogEventUpdateOne) ExecX(ctx context.Context) {
	if err := leuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (leuo *LogEventUpdateOne) sqlSave(ctx context.Context) (_node *LogEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   logevent.Table,
			Columns: logevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logevent.FieldID,
			},
		},
	}
	id, ok := leuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LogEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := leuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logevent.FieldID)
		for _, f := range fields {
			if !logevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != logevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := leuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := leuo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: logevent.FieldTimestamp,
		})
	}
	if value, ok := leuo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: logevent.FieldLevel,
		})
	}
	if value, ok := leuo.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: logevent.FieldLevel,
		})
	}
	if value, ok := leuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logevent.FieldMessage,
		})
	}
	_node = &LogEvent{config: leuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, leuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
