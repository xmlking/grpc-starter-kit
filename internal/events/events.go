package events

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog/log"
	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/ent/hook"
)

func ProfileChangeEvents() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ProfileFunc(func(ctx context.Context, m *ent.ProfileMutation) (ent.Value, error) {
			avatarURL, exists := m.Avatar()
			if !exists {
				return nil, errors.New("avatar field is missing")
			}
			log.Debug().Msgf("avatarURL:%s", avatarURL)
			// TODO:
			// 1. Verify that "avatarURL" points to a real object in the bucket.
			// 2. Otherwise, fail.
			return next.Mutate(ctx, m)
		})
	}
	// Limit the hook only to "Create" and "Update" operations.
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
