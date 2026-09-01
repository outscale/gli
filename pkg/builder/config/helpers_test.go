package configbuilder_test

import (
	"testing"

	configbuilder "github.com/outscale/octl/pkg/builder/config"
	"github.com/stretchr/testify/assert"
)

func TestSingular(t *testing.T) {
	tts := []struct {
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
		assert.Equal(t, tt.singular, configbuilder.Singular(tt.plural))
	}
}
