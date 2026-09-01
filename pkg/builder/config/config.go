package configbuilder

import (
	"strings"

	"github.com/outscale/octl/pkg/config"
	"github.com/samber/lo"
)

type Config struct {
	InputStructSuffix string

	// Methods

	// SkipMethods skips the method (no octl API command)
	SkipMethods []string
	// SkipSuffixes skips the method with the following suffixes
	SkipSuffixes []string
	// AllowedNumOut defines the number of allowed args
	AllowedNumOut []int

	// Entities

	// Skip avoids generating an alias for an entity
	SkipAlias []string
	// TypeName is an entity => type name map
	TypeName map[string]string

	// Fields

	PriorityFields           []string
	RequiredFromFieldPointer bool
	RequiredFromComment      func(string) bool

	// Flags

	// ReadFlagPrefixes The prefix that read flags must have (removed in alias flag).
	ReadFlagPrefixes []string
	// CreateFlagPrefixes The prefix that create flags must have (removed in alias flag).
	CreateFlagPrefixes []string
	// UpdateFlagPrefixes The prefix that update flags must have (removed in alias flag).
	UpdateFlagPrefixes []string
	// DeleteFlagPrefixes The prefix that delete flags must have (removed in alias flag).
	DeleteFlagPrefixes []string

	SkipFlagsPrefixes  map[string][]string
	StripFlagsPrefixes map[string][]string

	FlagOverrides map[string]config.Flag
	FlagReplaces  []string

	IdFieldInUsage string

	AliasRootPath []string
}

func (c *Config) getEntityNameFor(method string) string {
	return strings.ToLower(c.getRawTypeNameFor(method))
}

func (c *Config) getTypeNameFor(method string) string {
	raw := c.getRawTypeNameFor(method)
	if t, found := c.TypeName[raw]; found {
		return t
	}
	return raw
}

func (c *Config) getRawTypeNameFor(method string) string {
	typesName := method
	words := lo.Words(typesName)
	if len(words) > 1 {
		typesName = strings.TrimPrefix(typesName, words[0])
	}
	typesName = strings.TrimSuffix(typesName, "V2")
	return Singular(typesName)
}

func (c *Config) skipFlagsPrefixes(entity string) []string {
	return append(c.SkipFlagsPrefixes["*"], c.SkipFlagsPrefixes[entity]...)
}

func (c *Config) stripFlagsPrefixes(entity string) []string {
	return append(c.StripFlagsPrefixes["*"], c.StripFlagsPrefixes[entity]...)
}
