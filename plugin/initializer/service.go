package initializer

import "context"

type PluginInitializer interface {
	Init(ctx context.Context, nodeIdentity string, rawConfiguration []byte) error
}
