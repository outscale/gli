package watch

import (
	"context"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/outscale/octl/pkg/output/format"
	"github.com/outscale/octl/pkg/spinner"
)

type RefreshMsg struct{}

// Run runs the watcher.
func Run[Error error](ctx context.Context, fmter format.Interface, fn func(ctx context.Context) error, interval time.Duration) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	model, tui := fmter.(*Format)
	var prog *tea.Program
	if tui {
		prog = tea.NewProgram(model)
		cancel = spinner.Run(ctx, fmt.Sprintf("Refreshing every %s... Press ctrl+c to exit", interval), false)
		defer cancel()
	}

	// Start the API calling loop
	exitErr := make(chan error, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := fn(ctx)
				if err != nil {
					exitErr <- err
					if prog != nil {
						prog.Quit()
					}
					return
				}
				if prog != nil {
					// prog.Send(RefreshMsg{})
					// The cursor sometimes reappears... This also forces a refresh.
					prog.Send(tea.HideCursor())
				}
				time.Sleep(interval)
			}
		}
	}()

	if tui {
		// Start Bubbletea
		if _, err := prog.Run(); err != nil {
			return fmt.Errorf("watch: %w", err)
		}
	} else {
		_ = spinner.Run(ctx, fmt.Sprintf("Refreshing every %s... Press ctrl+c to exit", interval), true)
	}

	// Forward the error from the API loop
	select {
	case err := <-exitErr:
		return err
	default:
		return nil
	}
}
