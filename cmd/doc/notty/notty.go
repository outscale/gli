package notty

import (
	"os"

	"github.com/outscale/octl/pkg/messages"
)

// trick the command help not to be rendered in color
// this init must be executed before the cmd init, generating the command help messages
func init() {
	fd, err := os.CreateTemp("", "doc")
	if err != nil {
		messages.ExitErr(err)
	}
	os.Stdout = fd

	// lipgloss.SetDefaultRenderer(lipgloss.NewRenderer(fd))
}
