// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rgst-io/jsonnet-playground/ent/code"
)

// CodeCreate is the builder for creating a Code entity.
type CodeCreate struct {
	config
	mutation *CodeMutation
	hooks    []Hook
}

// SetContents sets the "contents" field.
func (cc *CodeCreate) SetContents(s string) *CodeCreate {
	cc.mutation.SetContents(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CodeCreate) SetID(u uuid.UUID) *CodeCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CodeCreate) SetNillableID(u *uuid.UUID) *CodeCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the CodeMutation object of the builder.
func (cc *CodeCreate) Mutation() *CodeMutation {
	return cc.mutation
}

// Save creates the Code in the database.
func (cc *CodeCreate) Save(ctx context.Context) (*Code, error) {
	var (
		err  error
		node *Code
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CodeCreate) SaveX(ctx context.Context) *Code {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CodeCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CodeCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CodeCreate) defaults() {
	if _, ok := cc.mutation.ID(); !ok {
		v := code.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CodeCreate) check() error {
	if _, ok := cc.mutation.Contents(); !ok {
		return &ValidationError{Name: "contents", err: errors.New(`ent: missing required field "Code.contents"`)}
	}
	if v, ok := cc.mutation.Contents(); ok {
		if err := code.ContentsValidator(v); err != nil {
			return &ValidationError{Name: "contents", err: fmt.Errorf(`ent: validator failed for field "Code.contents": %w`, err)}
		}
	}
	return nil
}

func (cc *CodeCreate) sqlSave(ctx context.Context) (*Code, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cc *CodeCreate) createSpec() (*Code, *sqlgraph.CreateSpec) {
	var (
		_node = &Code{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: code.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: code.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Contents(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: code.FieldContents,
		})
		_node.Contents = value
	}
	return _node, _spec
}

// CodeCreateBulk is the builder for creating many Code entities in bulk.
type CodeCreateBulk struct {
	config
	builders []*CodeCreate
}

// Save creates the Code entities in the database.
func (ccb *CodeCreateBulk) Save(ctx context.Context) ([]*Code, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Code, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CodeMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CodeCreateBulk) SaveX(ctx context.Context) []*Code {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CodeCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
