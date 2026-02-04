package prerun

import (
	"context"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/outscale/gli/pkg/debug"
	"github.com/outscale/gli/pkg/errors"
	"github.com/outscale/gli/pkg/update"
	"github.com/outscale/gli/pkg/version"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

func CheckUpdate(cmd *cobra.Command, args []string) {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
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
	_, _ = color.New(color.FgGreen).Add(color.Bold).Printf("⬆️ New version %s detected - call gli update to update\n", latest)
}
