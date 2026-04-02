//go:build homebrew

package prerun

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/style"
	"github.com/outscale/octl/pkg/update"
	"github.com/outscale/octl/pkg/version"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

func CheckUpdate(cmd *cobra.Command, args []string) {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return
	}
	if no, _ := cmd.Flags().GetBool("no-upgrade"); no {
		return
	}
	ghCtx, cancel := context.WithTimeout(cmd.Context(), time.Second)
	defer cancel()
	latest := update.LatestRelease(ghCtx)
	debug.Println("check update", version.Version, latest, semver.Compare(version.Version, latest) < 0)
	if latest == "" || semver.Compare(version.Version, latest) >= 0 {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, style.Renderf(style.Yellow, `⬆️ New version %s detected - call "brew install octl" to update`, latest))
}
