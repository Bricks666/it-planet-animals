// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/predicate"
	"animals/ent/visitedlocation"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VisitedLocationDelete is the builder for deleting a VisitedLocation entity.
type VisitedLocationDelete struct {
	config
	hooks    []Hook
	mutation *VisitedLocationMutation
}

// Where appends a list predicates to the VisitedLocationDelete builder.
func (vld *VisitedLocationDelete) Where(ps ...predicate.VisitedLocation) *VisitedLocationDelete {
	vld.mutation.Where(ps...)
	return vld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vld *VisitedLocationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, VisitedLocationMutation](ctx, vld.sqlExec, vld.mutation, vld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vld *VisitedLocationDelete) ExecX(ctx context.Context) int {
	n, err := vld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vld *VisitedLocationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(visitedlocation.Table, sqlgraph.NewFieldSpec(visitedlocation.FieldID, field.TypeUint64))
	if ps := vld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vld.mutation.done = true
	return affected, err
}

// VisitedLocationDeleteOne is the builder for deleting a single VisitedLocation entity.
type VisitedLocationDeleteOne struct {
	vld *VisitedLocationDelete
}

// Where appends a list predicates to the VisitedLocationDelete builder.
func (vldo *VisitedLocationDeleteOne) Where(ps ...predicate.VisitedLocation) *VisitedLocationDeleteOne {
	vldo.vld.mutation.Where(ps...)
	return vldo
}

// Exec executes the deletion query.
func (vldo *VisitedLocationDeleteOne) Exec(ctx context.Context) error {
	n, err := vldo.vld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{visitedlocation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vldo *VisitedLocationDeleteOne) ExecX(ctx context.Context) {
	if err := vldo.Exec(ctx); err != nil {
		panic(err)
	}
}
