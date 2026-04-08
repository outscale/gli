package spinner

import (
	"context"
	"os"

	"charm.land/huh/v2/spinner"
	"github.com/outscale/octl/pkg/style"
)

func Run(ctx context.Context, text string) (cancel func()) {
	spinWait := make(chan struct{})
	spinCtx, spinCancel := context.WithCancel(ctx)

	theme := spinner.ThemeFunc(func(bool) *spinner.Styles {
		return &spinner.Styles{
			Spinner: style.Yellow,
			Title:   style.Faint,
		}
	})

	spin := spinner.New().
		Title(text).
		Context(spinCtx).
		WithOutput(os.Stderr).
		WithTheme(theme)
	go func() {
		_ = spin.Run()
		close(spinWait)
	}()

	return func() {
		spinCancel()
		<-spinWait
	}
}
