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
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/betsetting"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/predicate"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
)

// BetSettingUpdate is the builder for updating BetSetting entities.
type BetSettingUpdate struct {
	config
	hooks    []Hook
	mutation *BetSettingMutation
}

// Where appends a list predicates to the BetSettingUpdate builder.
func (bsu *BetSettingUpdate) Where(ps ...predicate.BetSetting) *BetSettingUpdate {
	bsu.mutation.Where(ps...)
	return bsu
}

// SetValues sets the "values" field.
func (bsu *BetSettingUpdate) SetValues(ssm *schema.BetSettingMap) *BetSettingUpdate {
	bsu.mutation.SetValues(ssm)
	return bsu
}

// SetCreatedTime sets the "created_time" field.
func (bsu *BetSettingUpdate) SetCreatedTime(t time.Time) *BetSettingUpdate {
	bsu.mutation.SetCreatedTime(t)
	return bsu
}

// AddSheetIDs adds the "sheets" edge to the Sheet entity by IDs.
func (bsu *BetSettingUpdate) AddSheetIDs(ids ...int64) *BetSettingUpdate {
	bsu.mutation.AddSheetIDs(ids...)
	return bsu
}

// AddSheets adds the "sheets" edges to the Sheet entity.
func (bsu *BetSettingUpdate) AddSheets(s ...*Sheet) *BetSettingUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bsu.AddSheetIDs(ids...)
}

// Mutation returns the BetSettingMutation object of the builder.
func (bsu *BetSettingUpdate) Mutation() *BetSettingMutation {
	return bsu.mutation
}

// ClearSheets clears all "sheets" edges to the Sheet entity.
func (bsu *BetSettingUpdate) ClearSheets() *BetSettingUpdate {
	bsu.mutation.ClearSheets()
	return bsu
}

// RemoveSheetIDs removes the "sheets" edge to Sheet entities by IDs.
func (bsu *BetSettingUpdate) RemoveSheetIDs(ids ...int64) *BetSettingUpdate {
	bsu.mutation.RemoveSheetIDs(ids...)
	return bsu
}

// RemoveSheets removes "sheets" edges to Sheet entities.
func (bsu *BetSettingUpdate) RemoveSheets(s ...*Sheet) *BetSettingUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bsu.RemoveSheetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bsu *BetSettingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BetSettingMutation](ctx, bsu.sqlSave, bsu.mutation, bsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bsu *BetSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := bsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bsu *BetSettingUpdate) Exec(ctx context.Context) error {
	_, err := bsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsu *BetSettingUpdate) ExecX(ctx context.Context) {
	if err := bsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bsu *BetSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   betsetting.Table,
			Columns: betsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: betsetting.FieldID,
			},
		},
	}
	if ps := bsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bsu.mutation.Values(); ok {
		_spec.SetField(betsetting.FieldValues, field.TypeJSON, value)
	}
	if value, ok := bsu.mutation.CreatedTime(); ok {
		_spec.SetField(betsetting.FieldCreatedTime, field.TypeTime, value)
	}
	if bsu.mutation.SheetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   betsetting.SheetsTable,
			Columns: []string{betsetting.SheetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sheet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsu.mutation.RemovedSheetsIDs(); len(nodes) > 0 && !bsu.mutation.SheetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   betsetting.SheetsTable,
			Columns: []string{betsetting.SheetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sheet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsu.mutation.SheetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   betsetting.SheetsTable,
			Columns: []string{betsetting.SheetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sheet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{betsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bsu.mutation.done = true
	return n, nil
}

// BetSettingUpdateOne is the builder for updating a single BetSetting entity.
type BetSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BetSettingMutation
}

// SetValues sets the "values" field.
func (bsuo *BetSettingUpdateOne) SetValues(ssm *schema.BetSettingMap) *BetSettingUpdateOne {
	bsuo.mutation.SetValues(ssm)
	return bsuo
}

// SetCreatedTime sets the "created_time" field.
func (bsuo *BetSettingUpdateOne) SetCreatedTime(t time.Time) *BetSettingUpdateOne {
	bsuo.mutation.SetCreatedTime(t)
	return bsuo
}

// AddSheetIDs adds the "sheets" edge to the Sheet entity by IDs.
func (bsuo *BetSettingUpdateOne) AddSheetIDs(ids ...int64) *BetSettingUpdateOne {
	bsuo.mutation.AddSheetIDs(ids...)
	return bsuo
}

// AddSheets adds the "sheets" edges to the Sheet entity.
func (bsuo *BetSettingUpdateOne) AddSheets(s ...*Sheet) *BetSettingUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bsuo.AddSheetIDs(ids...)
}

// Mutation returns the BetSettingMutation object of the builder.
func (bsuo *BetSettingUpdateOne) Mutation() *BetSettingMutation {
	return bsuo.mutation
}

// ClearSheets clears all "sheets" edges to the Sheet entity.
func (bsuo *BetSettingUpdateOne) ClearSheets() *BetSettingUpdateOne {
	bsuo.mutation.ClearSheets()
	return bsuo
}

// RemoveSheetIDs removes the "sheets" edge to Sheet entities by IDs.
func (bsuo *BetSettingUpdateOne) RemoveSheetIDs(ids ...int64) *BetSettingUpdateOne {
	bsuo.mutation.RemoveSheetIDs(ids...)
	return bsuo
}

// RemoveSheets removes "sheets" edges to Sheet entities.
func (bsuo *BetSettingUpdateOne) RemoveSheets(s ...*Sheet) *BetSettingUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bsuo.RemoveSheetIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bsuo *BetSettingUpdateOne) Select(field string, fields ...string) *BetSettingUpdateOne {
	bsuo.fields = append([]string{field}, fields...)
	return bsuo
}

// Save executes the query and returns the updated BetSetting entity.
func (bsuo *BetSettingUpdateOne) Save(ctx context.Context) (*BetSetting, error) {
	return withHooks[*BetSetting, BetSettingMutation](ctx, bsuo.sqlSave, bsuo.mutation, bsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bsuo *BetSettingUpdateOne) SaveX(ctx context.Context) *BetSetting {
	node, err := bsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bsuo *BetSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := bsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsuo *BetSettingUpdateOne) ExecX(ctx context.Context) {
	if err := bsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bsuo *BetSettingUpdateOne) sqlSave(ctx context.Context) (_node *BetSetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   betsetting.Table,
			Columns: betsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: betsetting.FieldID,
			},
		},
	}
	id, ok := bsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BetSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, betsetting.FieldID)
		for _, f := range fields {
			if !betsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != betsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bsuo.mutation.Values(); ok {
		_spec.SetField(betsetting.FieldValues, field.TypeJSON, value)
	}
	if value, ok := bsuo.mutation.CreatedTime(); ok {
		_spec.SetField(betsetting.FieldCreatedTime, field.TypeTime, value)
	}
	if bsuo.mutation.SheetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   betsetting.SheetsTable,
			Columns: []string{betsetting.SheetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sheet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsuo.mutation.RemovedSheetsIDs(); len(nodes) > 0 && !bsuo.mutation.SheetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   betsetting.SheetsTable,
			Columns: []string{betsetting.SheetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sheet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsuo.mutation.SheetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   betsetting.SheetsTable,
			Columns: []string{betsetting.SheetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sheet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BetSetting{config: bsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{betsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bsuo.mutation.done = true
	return _node, nil
}
