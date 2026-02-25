package descriptions

import (
	"regexp"
	"strings"
)

var reSpaces = regexp.MustCompile("[ ]{2,}")

func OneLine(desc string) string {
	r := strings.NewReplacer(
		"<br />", " ",
		"\\|", "|",
		"`", "",
		"\r", " ",
		"\n", " ",
		"&quot;", `"`,
	)
	return strings.TrimSpace(reSpaces.ReplaceAllString(r.Replace(desc), " "))
}

func Summary(desc string) string {
	description, _, found := strings.Cut(OneLine(desc), ". ")
	if found {
		description += "."
	}
	return description
}

func Clean(desc string) string {
	r := strings.NewReplacer(
		"<br />", `
`,
	)
	return strings.TrimSpace(reSpaces.ReplaceAllString(r.Replace(desc), " "))
}
