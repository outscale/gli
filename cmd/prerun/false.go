package prerun

import (
	"os"
	"slices"

	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/errors"
	"github.com/spf13/cobra"
)

func CheckFalse(cmd *cobra.Command, args []string) {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return
	}
	if slices.Contains(args, "false") {
		errors.Warn("--flag false does not work, use --flag=false")
	}
}
