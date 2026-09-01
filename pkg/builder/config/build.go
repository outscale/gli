package configbuilder

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/messages"
)

type MethodBuilder interface {
	Build(spec Spec) error
	Commit()
}

type ClientBuilder[T any] struct {
	cfg Config
}

func NewClientBuilder[T any](cfg Config) *ClientBuilder[T] {
	return &ClientBuilder[T]{
		cfg: cfg,
	}
}

func (b *ClientBuilder[T]) BuildFor(build *config.Config, spec Spec) {
	ct := reflect.TypeFor[T]()
	fmt.Printf("==== %s ====\n", ct.Name())
	// First, build read methods
	for m := range ct.Methods() {
		if isReadMethod(m.Name) {
			b.BuildMethod(build, m, spec)
		}
	}
	// Next, build write methods (some write aliases may require a read)
	for m := range ct.Methods() {
		if !isReadMethod(m.Name) {
			b.BuildMethod(build, m, spec)
		}
	}
}

func (b *ClientBuilder[T]) BuildMethod(build *config.Config, m reflect.Method, spec Spec) {
	if slices.Contains(b.cfg.SkipMethods, m.Name) {
		fmt.Println("***", m.Name, "(skipped)")
		return
	}
	if slices.ContainsFunc(b.cfg.SkipSuffixes, func(suf string) bool { return strings.HasSuffix(m.Name, suf) }) {
		fmt.Println("***", m.Name, "(skipped/suffix)")
		return
	}
	if !slices.Contains(b.cfg.AllowedNumOut, m.Type.NumOut()) {
		fmt.Println("***", m.Name, "(skipped/numout)")
		return
	}
	fmt.Println("***", m.Name)
	callMethodBuilder(NewAPIBuilder(b.cfg, build, m), spec)
	callMethodBuilder(NewEntityBuilder(b.cfg, build, m), spec)
	callMethodBuilder(NewAliasBuilder(b.cfg, build, m), spec)
}

func callMethodBuilder(mb MethodBuilder, spec Spec) {
	err := mb.Build(spec)
	switch {
	case errors.Is(err, ErrCantBuild):
	case err != nil:
		messages.ExitErr(err)
	default:
		mb.Commit()
	}
}
