package errors

import (
	"os"

	"github.com/fatih/color"
)

func Info(format string, a ...any) {
	_, _ = color.New(color.FgYellow).Add(color.Faint).Fprintf(os.Stderr, format+"\n", a...)
}

func Warn(format string, a ...any) {
	_, _ = color.New(color.FgYellow).Fprintf(os.Stderr, format+"\n", a...)
}
