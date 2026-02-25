package builder_test

import (
	"testing"

	"github.com/outscale/octl/pkg/config/generate/builder"
	"github.com/stretchr/testify/assert"
)

func TestSingular(t *testing.T) {
	var tts = []struct {
		plural, singular string
	}{
		{"Ips", "Ip"},
		{"Dns", "dns"},
		{"dns", "dns"},
		{"iops", "iops"},
		{"is", "is"},
		{"entities", "entity"},
		{"entries", "entry"},
		{"objects", "object"},
		{"nets", "net"},
		{"flexiblegpus", "flexiblegpu"},
		{"data", "data"},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.singular, builder.Singular(tt.plural))
	}
}
