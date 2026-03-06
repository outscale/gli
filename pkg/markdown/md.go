package markdown

import (
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/x/term"
	"github.com/mattn/go-isatty"
)

type Renderer interface {
	Render(string) (string, error)
}

func NewRenderer() Renderer {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return baseRenderer{}
	}
	termWidth, _, _ := term.GetSize(os.Stdout.Fd())
	if termWidth > 0 {
		termWidth = min(120, termWidth)
	}
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(termWidth),
	)
	if err != nil {
		return baseRenderer{}
	}
	return r
}

type baseRenderer struct{}

var reEOL = regexp.MustCompile("[\r\n]{2,}")

func (baseRenderer) Render(str string) (string, error) {
	r := strings.NewReplacer(
		"<br />", "\n",
		"\\|", "|",
		"\r\n", "\n",
	)
	return reEOL.ReplaceAllString(r.Replace(str), "\n\n"), nil
}
