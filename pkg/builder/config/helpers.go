package configbuilder

import (
	"slices"
	"strings"

	"github.com/gobuffalo/flect"
)

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

var (
	readPrefixes  = []string{"Read", "List", "Get"}
	writePrefixes = []string{"Delete", "Update", "Put", "Create"}
)

func isReadMethod(method string) bool {
	return slices.ContainsFunc(readPrefixes, func(pre string) bool { return strings.HasPrefix(method, pre) })
}

func isWriteMethod(method string) bool {
	return slices.ContainsFunc(writePrefixes, func(pre string) bool { return strings.HasPrefix(method, pre) })
}
