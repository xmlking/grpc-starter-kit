package schema

import (
	"net/url"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Annotations for codegen
func (Profile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entproto.Message(),
		// entproto.Service(),
	}
}

// Mixin for codegen
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
			Immutable().
			Annotations(entproto.Field(1)),
		field.Int("age").
			Positive().
			Annotations(entproto.Field(2)),
		field.String("tz"). // *time.Location?
					Annotations(entproto.Field(3)),
		field.JSON("avatar", &url.URL{}).
			Optional().
			Annotations(entproto.Field(4)),
		field.Time("birthday").
			Optional().
			Annotations(entproto.Field(5)),
		field.Enum("gender").
			Values("male", "female", "unspecified").
			Optional().
			Annotations(
				entproto.Field(6),
				entproto.Enum(map[string]int32{
					"male":        1,
					"female":      2,
					"unspecified": 3,
				}),
			),
		field.String("preferred_theme").
			Optional().
			Annotations(entproto.Field(7)),
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
