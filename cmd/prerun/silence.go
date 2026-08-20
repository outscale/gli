package prerun

import (
	"io"

	"github.com/outscale/octl/pkg/messages"
	"github.com/spf13/cobra"
)

func Silence(cmd *cobra.Command) {
	silent, _ := cmd.Flags().GetBool("silent")
	if silent {
		messages.MsgOut = io.Discard
	}
}
