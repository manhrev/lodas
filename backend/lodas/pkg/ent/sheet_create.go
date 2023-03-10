// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/betsetting"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/record"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
)

// SheetCreate is the builder for creating a Sheet entity.
type SheetCreate struct {
	config
	mutation *SheetMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (sc *SheetCreate) SetStatus(i int64) *SheetCreate {
	sc.mutation.SetStatus(i)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *SheetCreate) SetNillableStatus(i *int64) *SheetCreate {
	if i != nil {
		sc.SetStatus(*i)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SheetCreate) SetName(s string) *SheetCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetArea sets the "area" field.
func (sc *SheetCreate) SetArea(i int64) *SheetCreate {
	sc.mutation.SetArea(i)
	return sc
}

// SetNillableArea sets the "area" field if the given value is not nil.
func (sc *SheetCreate) SetNillableArea(i *int64) *SheetCreate {
	if i != nil {
		sc.SetArea(*i)
	}
	return sc
}

// SetProvince sets the "province" field.
func (sc *SheetCreate) SetProvince(i int64) *SheetCreate {
	sc.mutation.SetProvince(i)
	return sc
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (sc *SheetCreate) SetNillableProvince(i *int64) *SheetCreate {
	if i != nil {
		sc.SetProvince(*i)
	}
	return sc
}

// SetRatio sets the "ratio" field.
func (sc *SheetCreate) SetRatio(f float64) *SheetCreate {
	sc.mutation.SetRatio(f)
	return sc
}

// SetNillableRatio sets the "ratio" field if the given value is not nil.
func (sc *SheetCreate) SetNillableRatio(f *float64) *SheetCreate {
	if f != nil {
		sc.SetRatio(*f)
	}
	return sc
}

// SetWinRatio sets the "win_ratio" field.
func (sc *SheetCreate) SetWinRatio(f float64) *SheetCreate {
	sc.mutation.SetWinRatio(f)
	return sc
}

// SetNillableWinRatio sets the "win_ratio" field if the given value is not nil.
func (sc *SheetCreate) SetNillableWinRatio(f *float64) *SheetCreate {
	if f != nil {
		sc.SetWinRatio(*f)
	}
	return sc
}

// SetTotalCashin sets the "total_cashin" field.
func (sc *SheetCreate) SetTotalCashin(i int64) *SheetCreate {
	sc.mutation.SetTotalCashin(i)
	return sc
}

// SetNillableTotalCashin sets the "total_cashin" field if the given value is not nil.
func (sc *SheetCreate) SetNillableTotalCashin(i *int64) *SheetCreate {
	if i != nil {
		sc.SetTotalCashin(*i)
	}
	return sc
}

// SetTotalCashout sets the "total_cashout" field.
func (sc *SheetCreate) SetTotalCashout(i int64) *SheetCreate {
	sc.mutation.SetTotalCashout(i)
	return sc
}

// SetNillableTotalCashout sets the "total_cashout" field if the given value is not nil.
func (sc *SheetCreate) SetNillableTotalCashout(i *int64) *SheetCreate {
	if i != nil {
		sc.SetTotalCashout(*i)
	}
	return sc
}

// SetResultTime sets the "result_time" field.
func (sc *SheetCreate) SetResultTime(t time.Time) *SheetCreate {
	sc.mutation.SetResultTime(t)
	return sc
}

// SetCreatedTime sets the "created_time" field.
func (sc *SheetCreate) SetCreatedTime(t time.Time) *SheetCreate {
	sc.mutation.SetCreatedTime(t)
	return sc
}

// SetUpdatedTime sets the "updated_time" field.
func (sc *SheetCreate) SetUpdatedTime(t time.Time) *SheetCreate {
	sc.mutation.SetUpdatedTime(t)
	return sc
}

// SetUserID sets the "user_id" field.
func (sc *SheetCreate) SetUserID(i int64) *SheetCreate {
	sc.mutation.SetUserID(i)
	return sc
}

// SetID sets the "id" field.
func (sc *SheetCreate) SetID(i int64) *SheetCreate {
	sc.mutation.SetID(i)
	return sc
}

// AddRecordIDs adds the "records" edge to the Record entity by IDs.
func (sc *SheetCreate) AddRecordIDs(ids ...int64) *SheetCreate {
	sc.mutation.AddRecordIDs(ids...)
	return sc
}

// AddRecords adds the "records" edges to the Record entity.
func (sc *SheetCreate) AddRecords(r ...*Record) *SheetCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sc.AddRecordIDs(ids...)
}

// SetBetSettingID sets the "bet_setting" edge to the BetSetting entity by ID.
func (sc *SheetCreate) SetBetSettingID(id int64) *SheetCreate {
	sc.mutation.SetBetSettingID(id)
	return sc
}

// SetNillableBetSettingID sets the "bet_setting" edge to the BetSetting entity by ID if the given value is not nil.
func (sc *SheetCreate) SetNillableBetSettingID(id *int64) *SheetCreate {
	if id != nil {
		sc = sc.SetBetSettingID(*id)
	}
	return sc
}

// SetBetSetting sets the "bet_setting" edge to the BetSetting entity.
func (sc *SheetCreate) SetBetSetting(b *BetSetting) *SheetCreate {
	return sc.SetBetSettingID(b.ID)
}

// Mutation returns the SheetMutation object of the builder.
func (sc *SheetCreate) Mutation() *SheetMutation {
	return sc.mutation
}

// Save creates the Sheet in the database.
func (sc *SheetCreate) Save(ctx context.Context) (*Sheet, error) {
	sc.defaults()
	return withHooks[*Sheet, SheetMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SheetCreate) SaveX(ctx context.Context) *Sheet {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SheetCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SheetCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SheetCreate) defaults() {
	if _, ok := sc.mutation.Status(); !ok {
		v := sheet.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.Area(); !ok {
		v := sheet.DefaultArea
		sc.mutation.SetArea(v)
	}
	if _, ok := sc.mutation.Province(); !ok {
		v := sheet.DefaultProvince
		sc.mutation.SetProvince(v)
	}
	if _, ok := sc.mutation.Ratio(); !ok {
		v := sheet.DefaultRatio
		sc.mutation.SetRatio(v)
	}
	if _, ok := sc.mutation.WinRatio(); !ok {
		v := sheet.DefaultWinRatio
		sc.mutation.SetWinRatio(v)
	}
	if _, ok := sc.mutation.TotalCashin(); !ok {
		v := sheet.DefaultTotalCashin
		sc.mutation.SetTotalCashin(v)
	}
	if _, ok := sc.mutation.TotalCashout(); !ok {
		v := sheet.DefaultTotalCashout
		sc.mutation.SetTotalCashout(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SheetCreate) check() error {
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Sheet.status"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Sheet.name"`)}
	}
	if _, ok := sc.mutation.Area(); !ok {
		return &ValidationError{Name: "area", err: errors.New(`ent: missing required field "Sheet.area"`)}
	}
	if _, ok := sc.mutation.Province(); !ok {
		return &ValidationError{Name: "province", err: errors.New(`ent: missing required field "Sheet.province"`)}
	}
	if _, ok := sc.mutation.Ratio(); !ok {
		return &ValidationError{Name: "ratio", err: errors.New(`ent: missing required field "Sheet.ratio"`)}
	}
	if _, ok := sc.mutation.WinRatio(); !ok {
		return &ValidationError{Name: "win_ratio", err: errors.New(`ent: missing required field "Sheet.win_ratio"`)}
	}
	if _, ok := sc.mutation.TotalCashin(); !ok {
		return &ValidationError{Name: "total_cashin", err: errors.New(`ent: missing required field "Sheet.total_cashin"`)}
	}
	if _, ok := sc.mutation.TotalCashout(); !ok {
		return &ValidationError{Name: "total_cashout", err: errors.New(`ent: missing required field "Sheet.total_cashout"`)}
	}
	if _, ok := sc.mutation.ResultTime(); !ok {
		return &ValidationError{Name: "result_time", err: errors.New(`ent: missing required field "Sheet.result_time"`)}
	}
	if _, ok := sc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "Sheet.created_time"`)}
	}
	if _, ok := sc.mutation.UpdatedTime(); !ok {
		return &ValidationError{Name: "updated_time", err: errors.New(`ent: missing required field "Sheet.updated_time"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Sheet.user_id"`)}
	}
	return nil
}

func (sc *SheetCreate) sqlSave(ctx context.Context) (*Sheet, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SheetCreate) createSpec() (*Sheet, *sqlgraph.CreateSpec) {
	var (
		_node = &Sheet{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sheet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sheet.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(sheet.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(sheet.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Area(); ok {
		_spec.SetField(sheet.FieldArea, field.TypeInt64, value)
		_node.Area = value
	}
	if value, ok := sc.mutation.Province(); ok {
		_spec.SetField(sheet.FieldProvince, field.TypeInt64, value)
		_node.Province = value
	}
	if value, ok := sc.mutation.Ratio(); ok {
		_spec.SetField(sheet.FieldRatio, field.TypeFloat64, value)
		_node.Ratio = value
	}
	if value, ok := sc.mutation.WinRatio(); ok {
		_spec.SetField(sheet.FieldWinRatio, field.TypeFloat64, value)
		_node.WinRatio = value
	}
	if value, ok := sc.mutation.TotalCashin(); ok {
		_spec.SetField(sheet.FieldTotalCashin, field.TypeInt64, value)
		_node.TotalCashin = value
	}
	if value, ok := sc.mutation.TotalCashout(); ok {
		_spec.SetField(sheet.FieldTotalCashout, field.TypeInt64, value)
		_node.TotalCashout = value
	}
	if value, ok := sc.mutation.ResultTime(); ok {
		_spec.SetField(sheet.FieldResultTime, field.TypeTime, value)
		_node.ResultTime = value
	}
	if value, ok := sc.mutation.CreatedTime(); ok {
		_spec.SetField(sheet.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = value
	}
	if value, ok := sc.mutation.UpdatedTime(); ok {
		_spec.SetField(sheet.FieldUpdatedTime, field.TypeTime, value)
		_node.UpdatedTime = value
	}
	if value, ok := sc.mutation.UserID(); ok {
		_spec.SetField(sheet.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if nodes := sc.mutation.RecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sheet.RecordsTable,
			Columns: []string{sheet.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: record.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.BetSettingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sheet.BetSettingTable,
			Columns: []string{sheet.BetSettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: betsetting.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.bet_setting_sheets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SheetCreateBulk is the builder for creating many Sheet entities in bulk.
type SheetCreateBulk struct {
	config
	builders []*SheetCreate
}

// Save creates the Sheet entities in the database.
func (scb *SheetCreateBulk) Save(ctx context.Context) ([]*Sheet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sheet, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SheetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SheetCreateBulk) SaveX(ctx context.Context) []*Sheet {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SheetCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SheetCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
