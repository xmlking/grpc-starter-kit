package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/mixin"
)

type DeleteMixin struct {
	mixin.Schema
}

func (DeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Default(time.Now).
			Immutable(),
	}
}
