package cmd

import (
	"github.com/charmbracelet/huh"
	"github.com/outscale/octl/pkg/style"
)

func Prompt(question string, mode ...huh.EchoMode) (string, error) {
	echo := huh.EchoModeNormal
	if len(mode) > 0 {
		echo = mode[0]
	}
	var resp string
	err := huh.NewInput().
		Title(question).
		Prompt(">").
		EchoMode(echo).
		Value(&resp).
		WithTheme(style.Theme()).
		Run()
	return resp, err
}
