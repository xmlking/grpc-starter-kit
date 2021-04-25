package schema

import (
	"regexp"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations for codegen
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// FIXME: https://github.com/ent/ent/discussions/1433
		// entproto.Message(),
		// entproto.Service(),
	}
}

// Mixin for codegen
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		DeleteMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New).
			Immutable().
			Annotations(entproto.Field(1)),
		field.String("username").
			Unique().
			Immutable().
			MinLen(4).
			MaxLen(25).
			Annotations(entproto.Field(2)),
		field.String("first_name").
			NotEmpty().
			Annotations(entproto.Field(3)),
		field.String("last_name").
			NotEmpty().
			Annotations(entproto.Field(4)),
		field.String("email").
			Unique().
			NotEmpty().
			Match(emailRegex).
			MinLen(3).
			MaxLen(64).
			Annotations(entproto.Field(5)),
		field.String("tenant").
			StorageKey("organization").
			Default("demo").
			Immutable().
			Annotations(entproto.Field(6)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Unique().
			Annotations(entproto.Field(7)),
	}
}

// Indexes of user entity.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email", "tenant").
			Unique(),
		index.Fields("delete_time"),
	}
}
