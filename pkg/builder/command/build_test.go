package commandbuilder_test

import (
	"testing"

	commandbuilder "github.com/outscale/octl/pkg/builder/command"
	"github.com/spf13/cobra"
)

func BenchmarkBuild(b *testing.B) {
	for b.Loop() {
		cmd := &cobra.Command{
			Use: "cmd",
		}
		b := commandbuilder.NewBuilder("iaas", "https://docs.outscale.com/api.html")
		b.Build(cmd, nil)
	}
}
