// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/xmlking/grpc-starter-kit/ent/predicate"
	"github.com/xmlking/grpc-starter-kit/ent/profile"
	"github.com/xmlking/grpc-starter-kit/ent/user"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks      []Hook
	mutation   *ProfileMutation
	predicates []predicate.Profile
}

// Where adds a new predicate for the builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetAge sets the age field.
func (pu *ProfileUpdate) SetAge(i int) *ProfileUpdate {
	pu.mutation.ResetAge()
	pu.mutation.SetAge(i)
	return pu
}

// AddAge adds i to age.
func (pu *ProfileUpdate) AddAge(i int) *ProfileUpdate {
	pu.mutation.AddAge(i)
	return pu
}

// SetUsername sets the username field.
func (pu *ProfileUpdate) SetUsername(s string) *ProfileUpdate {
	pu.mutation.SetUsername(s)
	return pu
}

// SetTz sets the tz field.
func (pu *ProfileUpdate) SetTz(s string) *ProfileUpdate {
	pu.mutation.SetTz(s)
	return pu
}

// SetAvatar sets the avatar field.
func (pu *ProfileUpdate) SetAvatar(u *url.URL) *ProfileUpdate {
	pu.mutation.SetAvatar(u)
	return pu
}

// ClearAvatar clears the value of avatar.
func (pu *ProfileUpdate) ClearAvatar() *ProfileUpdate {
	pu.mutation.ClearAvatar()
	return pu
}

// SetBirthday sets the birthday field.
func (pu *ProfileUpdate) SetBirthday(t time.Time) *ProfileUpdate {
	pu.mutation.SetBirthday(t)
	return pu
}

// SetNillableBirthday sets the birthday field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableBirthday(t *time.Time) *ProfileUpdate {
	if t != nil {
		pu.SetBirthday(*t)
	}
	return pu
}

// ClearBirthday clears the value of birthday.
func (pu *ProfileUpdate) ClearBirthday() *ProfileUpdate {
	pu.mutation.ClearBirthday()
	return pu
}

// SetGender sets the gender field.
func (pu *ProfileUpdate) SetGender(pr profile.Gender) *ProfileUpdate {
	pu.mutation.SetGender(pr)
	return pu
}

// SetNillableGender sets the gender field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableGender(pr *profile.Gender) *ProfileUpdate {
	if pr != nil {
		pu.SetGender(*pr)
	}
	return pu
}

// ClearGender clears the value of gender.
func (pu *ProfileUpdate) ClearGender() *ProfileUpdate {
	pu.mutation.ClearGender()
	return pu
}

// SetPreferredTheme sets the preferred_theme field.
func (pu *ProfileUpdate) SetPreferredTheme(s string) *ProfileUpdate {
	pu.mutation.SetPreferredTheme(s)
	return pu
}

// SetNillablePreferredTheme sets the preferred_theme field if the given value is not nil.
func (pu *ProfileUpdate) SetNillablePreferredTheme(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetPreferredTheme(*s)
	}
	return pu
}

// ClearPreferredTheme clears the value of preferred_theme.
func (pu *ProfileUpdate) ClearPreferredTheme() *ProfileUpdate {
	pu.mutation.ClearPreferredTheme()
	return pu
}

// SetUserID sets the user edge to User by id.
func (pu *ProfileUpdate) SetUserID(id uuid.UUID) *ProfileUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetUser sets the user edge to User.
func (pu *ProfileUpdate) SetUser(u *User) *ProfileUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// ClearUser clears the user edge to User.
func (pu *ProfileUpdate) ClearUser() *ProfileUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := profile.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
	if v, ok := pu.mutation.Age(); ok {
		if err := profile.AgeValidator(v); err != nil {
			return 0, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Gender(); ok {
		if err := profile.GenderValidator(v); err != nil {
			return 0, &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}

	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"user\"")
	}
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldAge,
		})
	}
	if value, ok := pu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldAge,
		})
	}
	if value, ok := pu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldUsername,
		})
	}
	if value, ok := pu.mutation.Tz(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldTz,
		})
	}
	if value, ok := pu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: profile.FieldAvatar,
		})
	}
	if pu.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: profile.FieldAvatar,
		})
	}
	if value, ok := pu.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldBirthday,
		})
	}
	if pu.mutation.BirthdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: profile.FieldBirthday,
		})
	}
	if value, ok := pu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: profile.FieldGender,
		})
	}
	if pu.mutation.GenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: profile.FieldGender,
		})
	}
	if value, ok := pu.mutation.PreferredTheme(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldPreferredTheme,
		})
	}
	if pu.mutation.PreferredThemeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldPreferredTheme,
		})
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProfileMutation
}

// SetAge sets the age field.
func (puo *ProfileUpdateOne) SetAge(i int) *ProfileUpdateOne {
	puo.mutation.ResetAge()
	puo.mutation.SetAge(i)
	return puo
}

// AddAge adds i to age.
func (puo *ProfileUpdateOne) AddAge(i int) *ProfileUpdateOne {
	puo.mutation.AddAge(i)
	return puo
}

// SetUsername sets the username field.
func (puo *ProfileUpdateOne) SetUsername(s string) *ProfileUpdateOne {
	puo.mutation.SetUsername(s)
	return puo
}

// SetTz sets the tz field.
func (puo *ProfileUpdateOne) SetTz(s string) *ProfileUpdateOne {
	puo.mutation.SetTz(s)
	return puo
}

// SetAvatar sets the avatar field.
func (puo *ProfileUpdateOne) SetAvatar(u *url.URL) *ProfileUpdateOne {
	puo.mutation.SetAvatar(u)
	return puo
}

// ClearAvatar clears the value of avatar.
func (puo *ProfileUpdateOne) ClearAvatar() *ProfileUpdateOne {
	puo.mutation.ClearAvatar()
	return puo
}

// SetBirthday sets the birthday field.
func (puo *ProfileUpdateOne) SetBirthday(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetBirthday(t)
	return puo
}

// SetNillableBirthday sets the birthday field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableBirthday(t *time.Time) *ProfileUpdateOne {
	if t != nil {
		puo.SetBirthday(*t)
	}
	return puo
}

// ClearBirthday clears the value of birthday.
func (puo *ProfileUpdateOne) ClearBirthday() *ProfileUpdateOne {
	puo.mutation.ClearBirthday()
	return puo
}

// SetGender sets the gender field.
func (puo *ProfileUpdateOne) SetGender(pr profile.Gender) *ProfileUpdateOne {
	puo.mutation.SetGender(pr)
	return puo
}

// SetNillableGender sets the gender field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableGender(pr *profile.Gender) *ProfileUpdateOne {
	if pr != nil {
		puo.SetGender(*pr)
	}
	return puo
}

// ClearGender clears the value of gender.
func (puo *ProfileUpdateOne) ClearGender() *ProfileUpdateOne {
	puo.mutation.ClearGender()
	return puo
}

// SetPreferredTheme sets the preferred_theme field.
func (puo *ProfileUpdateOne) SetPreferredTheme(s string) *ProfileUpdateOne {
	puo.mutation.SetPreferredTheme(s)
	return puo
}

// SetNillablePreferredTheme sets the preferred_theme field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillablePreferredTheme(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetPreferredTheme(*s)
	}
	return puo
}

// ClearPreferredTheme clears the value of preferred_theme.
func (puo *ProfileUpdateOne) ClearPreferredTheme() *ProfileUpdateOne {
	puo.mutation.ClearPreferredTheme()
	return puo
}

// SetUserID sets the user edge to User by id.
func (puo *ProfileUpdateOne) SetUserID(id uuid.UUID) *ProfileUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetUser sets the user edge to User.
func (puo *ProfileUpdateOne) SetUser(u *User) *ProfileUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// ClearUser clears the user edge to User.
func (puo *ProfileUpdateOne) ClearUser() *ProfileUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := profile.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
	if v, ok := puo.mutation.Age(); ok {
		if err := profile.AgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Gender(); ok {
		if err := profile.GenderValidator(v); err != nil {
			return nil, &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}

	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"user\"")
	}
	var (
		err  error
		node *Profile
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (pr *Profile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Profile.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldAge,
		})
	}
	if value, ok := puo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldAge,
		})
	}
	if value, ok := puo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldUsername,
		})
	}
	if value, ok := puo.mutation.Tz(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldTz,
		})
	}
	if value, ok := puo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: profile.FieldAvatar,
		})
	}
	if puo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: profile.FieldAvatar,
		})
	}
	if value, ok := puo.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldBirthday,
		})
	}
	if puo.mutation.BirthdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: profile.FieldBirthday,
		})
	}
	if value, ok := puo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: profile.FieldGender,
		})
	}
	if puo.mutation.GenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: profile.FieldGender,
		})
	}
	if value, ok := puo.mutation.PreferredTheme(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldPreferredTheme,
		})
	}
	if puo.mutation.PreferredThemeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldPreferredTheme,
		})
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Profile{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
