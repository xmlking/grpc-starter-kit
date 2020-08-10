// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/xmlking/grpc-starter-kit/ent/profile"
	"github.com/xmlking/grpc-starter-kit/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Tenant holds the value of the "tenant" field.
	Tenant string `json:"tenant,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Profile holds the value of the profile edge.
	Profile *Profile
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProfileOrErr() (*Profile, error) {
	if e.loadedTypes[0] {
		if e.Profile == nil {
			// The edge profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullTime{},   // delete_time
		&sql.NullString{}, // username
		&sql.NullString{}, // first_name
		&sql.NullString{}, // last_name
		&sql.NullString{}, // email
		&sql.NullString{}, // tenant
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		u.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		u.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		u.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field delete_time", values[2])
	} else if value.Valid {
		u.DeleteTime = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field username", values[3])
	} else if value.Valid {
		u.Username = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field first_name", values[4])
	} else if value.Valid {
		u.FirstName = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field last_name", values[5])
	} else if value.Valid {
		u.LastName = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[6])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tenant", values[7])
	} else if value.Valid {
		u.Tenant = value.String
	}
	return nil
}

// QueryProfile queries the profile edge of the User.
func (u *User) QueryProfile() *ProfileQuery {
	return (&UserClient{config: u.config}).QueryProfile(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(u.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(u.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(u.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", tenant=")
	builder.WriteString(u.Tenant)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
