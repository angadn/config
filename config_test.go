package config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/angadn/config"
)

func TestConfig(t *testing.T) {
	var (
		ctx = context.Background()
		cfg = config.FromEnv()
	)

	val, _ := cfg.Get(ctx, "PATH")
	def, _ := cfg.GetDef(ctx, "PATH", "path should be set")
	fmt.Printf("val = %s\ndef = %s\n", val, def)

	if def != val {
		t.Fail()
	}
}
