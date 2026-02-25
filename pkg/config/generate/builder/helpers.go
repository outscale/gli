package builder

import (
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/samber/lo"
)

func names(method string) (typeName, typesName, entityName string) {
	typesName = method
	words := lo.Words(typesName)
	if len(words) > 1 {
		typesName = strings.TrimPrefix(typesName, words[0])
	}
	typesName = strings.TrimSuffix(typesName, "V2")
	typeName = Singular(typesName)
	entityName = strings.ToLower(typeName)
	return
}

func Singular(s string) string {
	lc := strings.ToLower(s)
	switch lc {
	case "is", "iops", "options", "data", "details", "cors":
		return s
	case "publicips", "ips", "ids", "flexiblegpus":
		return strings.TrimSuffix(s, "s")
	}
	singular := flect.Singularize(s)
	// fix for flect bug
	if singular == s+strings.ToUpper(s) {
		return lc
	}
	return singular
}
