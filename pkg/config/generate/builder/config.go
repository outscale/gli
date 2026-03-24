package builder

import "github.com/outscale/octl/pkg/config"

type Config struct {
	InputStructSuffix string

	// ReadFlagPrefixes The prefix that read flags must have (removed in alias flag).
	ReadFlagPrefixes []string
	// CreateFlagPrefixes The prefix that create flags must have (removed in alias flag).
	CreateFlagPrefixes []string
	// UpdateFlagPrefixes The prefix that update flags must have (removed in alias flag).
	UpdateFlagPrefixes []string
	// DeleteFlagPrefixes The prefix that delete flags must have (removed in alias flag).
	DeleteFlagPrefixes []string

	SkipFlagsPrefixes []string

	FlagOverrides map[string]config.Flag
	FlagReplaces  []string

	PriorityFields []string

	RequiredFromFieldPointer bool
	RequiredFromComment      func(string) bool

	AliasRootPath []string
}
