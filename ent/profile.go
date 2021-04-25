// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/xmlking/grpc-starter-kit/ent/profile"
	"github.com/xmlking/grpc-starter-kit/ent/user"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Tz holds the value of the "tz" field.
	Tz string `json:"tz,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar *url.URL `json:"avatar,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday time.Time `json:"birthday,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender profile.Gender `json:"gender,omitempty"`
	// PreferredTheme holds the value of the "preferred_theme" field.
	PreferredTheme string `json:"preferred_theme,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProfileQuery when eager-loading is set.
	Edges        ProfileEdges `json:"edges"`
	user_profile *uuid.UUID
}

// ProfileEdges holds the relations/edges for other nodes in the graph.
type ProfileEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProfileEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case profile.FieldAvatar:
			values[i] = new([]byte)
		case profile.FieldAge:
			values[i] = new(sql.NullInt64)
		case profile.FieldTz, profile.FieldGender, profile.FieldPreferredTheme:
			values[i] = new(sql.NullString)
		case profile.FieldCreateTime, profile.FieldUpdateTime, profile.FieldDeleteTime, profile.FieldBirthday:
			values[i] = new(sql.NullTime)
		case profile.FieldID:
			values[i] = new(uuid.UUID)
		case profile.ForeignKeys[0]: // user_profile
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Profile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profile.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case profile.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pr.CreateTime = value.Time
			}
		case profile.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pr.UpdateTime = value.Time
			}
		case profile.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				pr.DeleteTime = new(time.Time)
				*pr.DeleteTime = value.Time
			}
		case profile.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				pr.Age = int(value.Int64)
			}
		case profile.FieldTz:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tz", values[i])
			} else if value.Valid {
				pr.Tz = value.String
			}
		case profile.FieldAvatar:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Avatar); err != nil {
					return fmt.Errorf("unmarshal field avatar: %w", err)
				}
			}
		case profile.FieldBirthday:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				pr.Birthday = value.Time
			}
		case profile.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pr.Gender = profile.Gender(value.String)
			}
		case profile.FieldPreferredTheme:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preferred_theme", values[i])
			} else if value.Valid {
				pr.PreferredTheme = value.String
			}
		case profile.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_profile", values[i])
			} else if value != nil {
				pr.user_profile = value
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Profile entity.
func (pr *Profile) QueryUser() *UserQuery {
	return (&ProfileClient{config: pr.config}).QueryUser(pr)
}

// Update returns a builder for updating this Profile.
// Note that you need to call Profile.Unwrap() before calling this method if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return (&ProfileClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Profile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profile is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(pr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(pr.UpdateTime.Format(time.ANSIC))
	if v := pr.DeleteTime; v != nil {
		builder.WriteString(", delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pr.Age))
	builder.WriteString(", tz=")
	builder.WriteString(pr.Tz)
	builder.WriteString(", avatar=")
	builder.WriteString(fmt.Sprintf("%v", pr.Avatar))
	builder.WriteString(", birthday=")
	builder.WriteString(pr.Birthday.Format(time.ANSIC))
	builder.WriteString(", gender=")
	builder.WriteString(fmt.Sprintf("%v", pr.Gender))
	builder.WriteString(", preferred_theme=")
	builder.WriteString(pr.PreferredTheme)
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile

func (pr Profiles) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
