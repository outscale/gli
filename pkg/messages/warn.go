package messages

import (
	"fmt"
	"os"

	"github.com/outscale/octl/pkg/style"
)

func Info(format string, a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, style.Renderf(style.Faint, format, a...))
}

func Warn(format string, a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, style.Renderf(style.Yellow, "⚠️ "+format, a...))
}

func Success(format string, a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, style.Renderf(style.Green, "✅ "+format, a...))
}
