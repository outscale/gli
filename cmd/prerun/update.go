package prerun

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/errors"
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
	ctx := context.Background()
	ghCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	latest, err := update.LatestRelease(ghCtx)
	if err != nil {
		errors.Info("❌ Unable to fetch latest release: %v", err)
		return
	}
	debug.Println("check update", version.Version, latest, semver.Compare(version.Version, latest) < 0)
	if latest == "" || semver.Compare(version.Version, latest) >= 0 {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, style.Renderf(style.Yellow, "⬆️ New version %s detected - call octl update to update", latest))
}
