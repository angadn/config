package config

import (
	"context"
	"os"
)

// Key is a type-alias to identify our Key strings.
type Key string

// Value is a type-alias to identify our Value strings.
type Value string

// FromEnv gives us a Source from the system's environment.
func FromEnv() (src Source) {
	src.SourceImpl = nilSourceImpl{}
	return
}

// Source for our configuration values.
type Source struct {
	SourceImpl
}

// Get the config Value for a Key from the local environment. If it is empty, check the
// underlying Source Implementation.
func (src Source) Get(ctx context.Context, key Key) (value Value, err error) {
	if value = Value(os.Getenv(string(key))); value != "" {
		return
	}

	value, err = src.SourceImpl.Get(ctx, key)
	return
}

// GetDef returns the value for a key if non-empty and no error, else the passed default
// value.
func (src Source) GetDef(
	ctx context.Context, key Key, def Value,
) (value Value, err error) {
	if value, err = src.SourceImpl.Get(ctx, key); err != nil || value == "" {
		value = def
	}

	return
}

// SourceImpl is an interface to implement for any configuration system.
type SourceImpl interface {
	Get(ctx context.Context, key Key) (value Value, err error)
	Set(ctx context.Context, key Key, value Value) (err error)
}

// nilSourceImpl always returns empty.
type nilSourceImpl struct{}

func (src nilSourceImpl) Get(
	ctx context.Context, key Key,
) (value Value, err error) {
	return
}

func (src nilSourceImpl) Set(
	ctx context.Context, key Key, value Value,
) (err error) {
	return
}
