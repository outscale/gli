package prerun

import (
	"os"
	"slices"

	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/messages"
	"github.com/spf13/cobra"
)

func CheckFalse(cmd *cobra.Command, args []string) {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return
	}
	if slices.Contains(args, "false") {
		messages.Warn("--flag false does not work, use --flag=false")
	}
}
