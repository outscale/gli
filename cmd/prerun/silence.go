package prerun

import (
	"github.com/outscale/octl/pkg/messages"
	"github.com/spf13/cobra"
)

func Silence(cmd *cobra.Command) {
	messages.Silent, _ = cmd.Flags().GetBool("silent")
}
