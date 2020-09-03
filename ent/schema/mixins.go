package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

type DeleteMixin struct {
	mixin.Schema
}

func (DeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Optional().
			Nillable(),
		//Default(time.Now).
		//Immutable(),
	}
}
