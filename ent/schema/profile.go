package schema

import (
	"net/url"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/facebook/ent/schema/mixin"
	"github.com/google/uuid"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

func (Profile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		DeleteMixin{},
	}
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New).
			Immutable(),
		field.Int("age").
			Positive(),
		field.String("tz"), // *time.Location?
		field.JSON("avatar", &url.URL{}).
			Optional(),
		field.Time("birthday").
			Optional(),
		field.Enum("gender").
			Values("male", "female", "unspecified").
			Optional(),
		field.String("preferred_theme").
			Optional(),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("profile").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Profile cannot be created without its owner.
			Required(),
	}
}

// Indexes of Profile entity.
func (Profile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("delete_time"),
	}
}
