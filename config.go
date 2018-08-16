package config

import "os"

// Provider can be set to configure which Source to use.
var Provider = func(args ...interface{}) (src Source) {
	src.SourceImpl = EnvSourceImpl{}
	return
}

// Source for our configuration values.
type Source struct {
	SourceImpl
}

// GetDef returns the value for a key if found, else the passed default value.
func (src Source) GetDef(key string, def string) (value string) {
	if value = src.SourceImpl.Get(key); len(value) > 0 {
		return
	}

	value = def
	return
}

// SourceImpl is an interface to implement for any configuration system.
type SourceImpl interface {
	Get(key string) (value string)
}

// EnvSourceImpl is an implementation of SourceImpl using os.Getenv.
type EnvSourceImpl struct{}

// Get a key from the environment.
func (src EnvSourceImpl) Get(key string) (value string) {
	return os.Getenv(key)
}
