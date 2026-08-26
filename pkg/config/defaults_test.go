package config_test

import (
	"testing"

	"github.com/outscale/octl/pkg/config"
)

func BenchmarkDefaults(b *testing.B) {
	for b.Loop() {
		_ = config.Defaults()
	}
}
