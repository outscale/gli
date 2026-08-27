package spinner

import (
	"context"
	"os"

	"github.com/charmbracelet/huh/spinner"
	"github.com/outscale/octl/pkg/style"
)

func Run(ctx context.Context, text string, foreground bool) (cancel func()) {
	spinWait := make(chan struct{})
	spinCtx, spinCancel := context.WithCancel(ctx)
	spin := spinner.New().
		Title(text).
		Context(spinCtx).
		Output(os.Stderr).
		Style(style.Yellow).
		TitleStyle(style.Faint)
	if foreground {
		_ = spin.Run()
		spinCancel()
		return
	}
	go func() {
		_ = spin.Run()
		close(spinWait)
	}()

	return func() {
		spinCancel()
		<-spinWait
	}
}
