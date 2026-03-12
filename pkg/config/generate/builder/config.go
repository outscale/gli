package builder

import "github.com/outscale/octl/pkg/config"

type Config struct {
	InputStructSuffix string

	// ReadFlagPrefix The prefix that read flags must have (removed in alias flag).
	ReadFlagPrefix string

	SkipFlags []string

	FlagOverrides map[string]config.Flag
	FlagReplaces  []string

	PriorityFields []string

	RequiredFromFieldPointer bool
	RequiredFromComment      func(string) bool
}
