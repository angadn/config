package config

import (
	"context"
	"os"
)

// FromEnv gives us a Source from the system's environment.
func FromEnv() (src Source) {
	src.SourceImpl = EnvSourceImpl{}
	return
}

// Source for our configuration values.
type Source struct {
	SourceImpl
}

// GetDef returns the value for a key if non-empty and no error, else the passed default
// value.
func (src Source) GetDef(
	ctx context.Context, key string, def string,
) (value string, err error) {
	if value, err = src.SourceImpl.Get(ctx, key); err != nil || value == "" {
		value = def
		return
	}

	return
}

// SourceImpl is an interface to implement for any configuration system.
type SourceImpl interface {
	Get(ctx context.Context, key string) (value string, err error)
}

// EnvSourceImpl is an implementation of SourceImpl using os.Getenv.
type EnvSourceImpl struct{}

// Get a key from the environment.
func (src EnvSourceImpl) Get(
	ctx context.Context, key string,
) (value string, err error) {
	value = os.Getenv(key)
	return
}
