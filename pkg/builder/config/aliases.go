package configbuilder

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"dario.cat/mergo"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/descriptions"
	"github.com/samber/lo"
)

var ErrCantBuild = errors.New("cannot build")

type AliasBuilder struct {
	cfg Config

	build *config.Config

	m reflect.Method

	typeName   string
	entityName string

	entity  config.Entity
	call    config.APICall
	aliases []config.Alias
}

func NewAliasBuilder(cfg Config, build *config.Config, m reflect.Method) *AliasBuilder {
	entity := cfg.getEntityNameFor(m.Name)
	typeName := cfg.getTypeNameFor(m.Name)
	call := build.API[m.Name]
	if call.Entity != "" {
		entity = call.Entity
	}
	return &AliasBuilder{
		cfg:        cfg,
		build:      build,
		m:          m,
		typeName:   typeName,
		entityName: entity,

		entity: build.Entities[entity],
		call:   call,
	}
}

func (b *AliasBuilder) Build(spec Spec) error {
	if slices.Contains(b.cfg.SkipAlias, b.m.Name) {
		return nil
	}
	if !isReadMethod(b.m.Name) && !isWriteMethod(b.m.Name) {
		return nil
	}

	fmt.Println(" - building alias")
	var err error
	switch {
	case strings.HasPrefix(b.m.Name, "Read"):
		err = b.buildListAlias()
		if err == nil {
			err = b.buildDescribeAlias()
		}
	case strings.HasPrefix(b.m.Name, "List"):
		err = b.buildListAlias()
	case strings.HasPrefix(b.m.Name, "Get"):
		err = b.buildDescribeAlias()
	case strings.HasPrefix(b.m.Name, "Delete"):
		err = b.buildDeleteAlias()
	case strings.HasPrefix(b.m.Name, "Update"):
		err = b.buildUpdateAlias("Update")
	case strings.HasPrefix(b.m.Name, "Put"):
		err = b.buildUpdateAlias("Put")
	case strings.HasPrefix(b.m.Name, "Create"):
		err = b.buildCreateAlias()
	default:
		return ErrCantBuild
	}
	return err
}

func normalizeFlag(cfg Config, prefixes []string) func(s string) string {
	return func(s string) string {
		for _, prefix := range prefixes {
			s = strings.TrimPrefix(s, prefix)
		}
		words := lo.Words(s)
		words = lo.FilterMap(words, func(s string, i int) (string, bool) {
			if s == "0" {
				return "", false
			}
			return Singular(strings.ToLower(s)), true
		})
		s = strings.Join(words, "-")
		if len(cfg.FlagReplaces) > 0 {
			rep := strings.NewReplacer(cfg.FlagReplaces...)
			s = rep.Replace(s)
		}
		words = strings.Split(s, "-")
		if len(words) > 4 && lo.HasSuffix(words[:len(words)-2], words[len(words)-2:]) {
			fmt.Println("  -> normalize reduce stutter", words, "=>", words[:len(words)-2])
			words = words[:len(words)-2]
		}
		if len(words) > 2 && lo.HasPrefix(words[1:], words[:1]) {
			fmt.Println("  -> normalize reduce stutter", words, "=>", words[1:])
			words = words[1:]
		}
		return strings.Join(words, "-")
	}
}

func (b *AliasBuilder) buildFlags(prefixes []string, ignore []string) config.FlagSet {
	ignore = append(ignore, b.cfg.skipFlagsPrefixes(b.entityName)...)

	normalize := normalizeFlag(b.cfg, prefixes)

	fs := b.build.API[b.m.Name].Flags
	cfs := config.FlagSet{}
	stripPrefixes := b.cfg.stripFlagsPrefixes(b.entityName)
	for _, f := range fs {
		if lo.ContainsBy(ignore, func(ignore string) bool { return ignore != "" && strings.HasPrefix(f.Name, ignore) }) {
			continue
		}
		flag := normalize(f.Name)
		if lo.ContainsBy(ignore, func(ignore string) bool { return ignore != "" && strings.HasPrefix(flag, ignore) }) {
			continue
		}
		stripped := strings.TrimPrefix(flag, lo.KebabCase(b.typeName)+"-")
		for _, prefix := range stripPrefixes {
			stripped = strings.TrimPrefix(flag, prefix)
		}
		if _, found := cfs.Get(stripped); !found {
			flag = stripped
		}

		cf := b.cfg.FlagOverrides[flag]
		fmt.Println("override for", flag, cf.Name)
		err := mergo.Merge(&cf, config.Flag{
			Name:     flag,
			AliasTo:  f.Name,
			Required: f.Required,
		})
		if err != nil {
			panic(err)
		}
		cfs = append(cfs, cf)
	}
	return cfs
}

func (b *AliasBuilder) requestIndex() int {
	reqIdx := 2
	if b.m.Type.In(0).Name() == "Context" {
		reqIdx = 1
	}
	return reqIdx
}

func (b *AliasBuilder) command(cmd ...string) []string {
	ncmd := slices.Clone(b.cfg.AliasRootPath)
	return append(ncmd, cmd...)
}

func (b *AliasBuilder) buildListAlias() error {
	reqIdx := b.requestIndex()
	if b.m.Type.NumIn() <= reqIdx {
		return nil
	}
	req := b.m.Type.In(reqIdx)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	var flags config.FlagSet
	if req.Kind() == reflect.Struct {
		flags = b.buildFlags(b.cfg.ReadFlagPrefixes, nil)
		fmt.Println("  -> list", b.typeName, "flags", flags.Names())
		b.aliases = append(b.aliases, config.Alias{
			Entity:    b.entityName,
			Use:       "list",
			Aliases:   []string{"ls"},
			AliasTo:   b.m.Name,
			AliasHelp: b.m.Name,
			Command: b.command(
				"api",
				b.m.Name,
				"--output", "table",
			),
			Flags: flags,
		})
	}
	return nil
}

func (b *AliasBuilder) buildDescribeAlias() error {
	reqIdx := b.requestIndex()
	if b.m.Type.NumIn() <= reqIdx {
		return nil
	}
	req := b.m.Type.In(reqIdx)
	var idFilter string
	if req.Kind() == reflect.Struct {
		// Guess id filter
		if b.entity.Primary == "" {
			return nil
		}
		flags := b.build.API[b.m.Name].Flags
		fids, found := lo.Find(flags, func(f config.Flag) bool {
			return slices.ContainsFunc(b.cfg.ReadFlagPrefixes, func(prefix string) bool { return strings.HasPrefix(f.Name, prefix+b.entity.Primary) })
		})
		if !found {
			fmt.Println("  -> no idFilter, no describe", b.typeName)
			return nil
		}
		idFilter = fids.Name
	}
	fmt.Println("  -> describe", b.typeName, "idFilter", idFilter)
	id := b.idFieldInUsage()
	cmd := []string{
		"api",
		b.m.Name,
	}
	if idFilter != "" {
		idFilter = "--" + idFilter
		cmd = append(cmd, idFilter)
	}
	cmd = append(cmd, "%*",
		"--output", "yaml",
		"--single",
	)
	b.aliases = append(b.aliases, config.Alias{
		Entity:    b.entityName,
		Use:       "describe " + id + " [" + id + "]...",
		Aliases:   []string{"desc"},
		AliasTo:   b.m.Name,
		AliasHelp: b.m.Name + " " + idFilter + " " + id,
		Command:   b.command(cmd...),
	})
	return nil
}

func (b *AliasBuilder) buildCreateAlias() error {
	flags := b.buildFlags(b.cfg.CreateFlagPrefixes, nil)
	fmt.Println("  -> create", b.typeName, "flags", flags.Names())
	b.aliases = append(b.aliases, config.Alias{
		Entity:    b.entityName,
		Use:       "create",
		Aliases:   []string{"add"},
		AliasTo:   b.m.Name,
		AliasHelp: b.m.Name,
		Command: b.command(
			"api",
			b.m.Name,
			"--output", "yaml",
			"--single",
		),
		Flags: flags,
	})
	return nil
}

func (b *AliasBuilder) buildUpdateAlias(verb string) error {
	idField, err := b.guessIDFilter()
	if err != nil {
		return err
	}
	verb = strings.ToLower(verb)
	reqIdx := b.requestIndex()
	if b.m.Type.NumIn() <= reqIdx {
		return nil
	}
	req := b.m.Type.In(reqIdx)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	if req.Kind() != reflect.Struct && b.m.Type.NumIn() >= reqIdx+2 {
		req = b.m.Type.In(reqIdx + 1)
	}
	var flags config.FlagSet
	if req.Kind() == reflect.Struct {
		flags = b.buildFlags(b.cfg.UpdateFlagPrefixes, []string{idField})
	}
	fmt.Println("  ->", verb, b.typeName, "idField", idField, "flags", flags.Names())
	id := b.idFieldInUsage()
	cmd := []string{
		"api",
		b.m.Name,
	}
	if idField != "" {
		idField = "--" + idField
		cmd = append(cmd, idField)
	}
	cmd = append(cmd, "%0", "--output", "yaml")

	b.aliases = append(b.aliases, config.Alias{
		Entity:    b.entityName,
		Use:       verb + " " + id + " [" + id + "]...",
		AliasTo:   b.m.Name,
		AliasHelp: b.m.Name + " " + idField + " " + id,
		Command:   b.command(cmd...),
		Flags:     flags,
	})
	return nil
}

func (b *AliasBuilder) guessIDFilter() (string, error) {
	req := b.m.Type.In(2)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	if req.Kind() != reflect.Struct {
		return "", nil
	}
	primary := b.build.Entities[b.entityName].Primary
	fids, found := req.FieldByName(primary)
	if !found {
		fids, found = req.FieldByName(primary + "s")
	}
	if !found {
		return "", ErrCantBuild
	}
	return fids.Name, nil
}

func (b *AliasBuilder) idFieldInUsage() string {
	if b.cfg.IdFieldInUsage != "" {
		return b.cfg.IdFieldInUsage
	}
	return lo.SnakeCase(b.entity.Primary)
}

func (b *AliasBuilder) buildDeleteAlias() error {
	idField, err := b.guessIDFilter()
	if err != nil {
		return err
	}
	reqIdx := b.requestIndex()
	if b.m.Type.NumIn() <= reqIdx {
		return nil
	}
	req := b.m.Type.In(reqIdx)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	if req.Kind() != reflect.Struct && b.m.Type.NumIn() >= reqIdx+2 {
		req = b.m.Type.In(reqIdx + 1)
	}
	var flags config.FlagSet
	if req.Kind() == reflect.Struct {
		flags = b.buildFlags(b.cfg.DeleteFlagPrefixes, []string{idField})
	}
	fmt.Println("  -> delete", b.typeName, "idField", idField, "flags", flags.Names())
	var displayCmd []string
	var displayFlags config.FlagSet
	for _, a := range b.build.Aliases {
		if a.Entity == b.entityName && strings.HasPrefix(a.Use, "describe") && a.SubCommand == "" {
			displayCmd = lo.Map(a.Command, func(arg string, i int) string {
				if i > 0 && (a.Command[i-1] == "-o" || a.Command[i-1] == "--output") {
					return "table"
				}
				return arg
			})
			displayFlags = lo.Filter(a.Flags, func(f config.Flag, _ int) bool {
				_, ok := flags.Get(f.Name)
				return ok
			})
		}
	}
	id := b.idFieldInUsage()
	cmd := []string{
		"api",
		b.m.Name,
	}
	if idField != "" {
		idField = "--" + idField
		cmd = append(cmd, idField)
	}
	cmd = append(cmd, "%0", "--output", "success")

	b.aliases = append(b.aliases, config.Alias{
		Entity:    b.entityName,
		Use:       "delete " + id + " [" + id + "]...",
		Aliases:   []string{"del", "rm"},
		AliasTo:   b.m.Name,
		AliasHelp: b.m.Name + " " + idField + " " + id,
		Command:   b.command(cmd...),
		Flags:     flags,
		Prompt: &config.Prompt{
			Action:         config.ActionDelete,
			DisplayCommand: displayCmd,
			Flags:          displayFlags,
		},
	})
	return nil
}

func (b *AliasBuilder) Commit() {
	for _, a := range b.aliases {
		if slices.ContainsFunc(b.build.Aliases, func(aa config.Alias) bool {
			return a.Entity == aa.Entity && a.Use == aa.Use
		}) {
			fmt.Println("  -> ### dropping duplicate", a.Use, a.Flags.Names())
			continue
		}
		b.build.Aliases = append(b.build.Aliases, a)
	}
	for i := range b.build.Aliases {
		a := &b.build.Aliases[i]
		if a.Help != "" && a.Short == "" {
			a.Short = descriptions.Summary(a.Help)
		}
	}
}
