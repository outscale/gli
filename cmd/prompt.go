package cmd

import (
	"github.com/charmbracelet/huh"
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
		Run()
	return resp, err
}
