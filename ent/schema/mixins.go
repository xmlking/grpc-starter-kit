package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// DeleteMixin struct
type DeleteMixin struct {
	mixin.Schema
}

// Fields method
func (DeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Optional().
			Nillable(),
		//Default(time.Now).
		//Immutable(),
	}
}
