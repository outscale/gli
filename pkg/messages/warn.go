package messages

import (
	"fmt"
	"io"
	"os"

	"github.com/outscale/octl/pkg/style"
)

var MsgOut io.Writer = os.Stderr

func Info(format string, a ...any) {
	_, _ = fmt.Fprintln(MsgOut, style.Renderf(style.Faint, format, a...))
}

func Err(format string, a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, style.Renderf(style.Error, "❌ "+format, a...))
}

func Warn(format string, a ...any) {
	_, _ = fmt.Fprintln(MsgOut, style.Renderf(style.Yellow, "⚠️ "+format, a...))
}

func Success(format string, a ...any) {
	_, _ = fmt.Fprintln(MsgOut, style.Renderf(style.Green, "✅ "+format, a...))
}
