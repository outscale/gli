package builder_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/outscale/octl/pkg/builder"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/spf13/cobra"
)

func BenchmarkBuild(b *testing.B) {
	for b.Loop() {
		cmd := &cobra.Command{
			Use: "cmd",
		}
		b := builder.NewBuilder[osc.Client]("iaas", "https://docs.outscale.com/api.html")
		b.BuildAPI(cmd, func(m reflect.Method) bool {
			return m.Type.NumIn() == 4 && m.Type.NumOut() == 2 && !strings.HasSuffix(m.Name, "Raw")
		}, nil)
		b.Build(cmd, nil)
	}
}
