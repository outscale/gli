package builder

import (
	"errors"
	"reflect"
	"strings"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/messages"
)

type ClientBuilder[T any] struct {
	cfg Config
}

func NewClientBuilder[T any](cfg Config) *ClientBuilder[T] {
	return &ClientBuilder[T]{
		cfg: cfg,
	}
}

func (b *ClientBuilder[T]) BuildFor(build *config.Config) {
	ct := reflect.TypeFor[T]()
	for m := range ct.Methods() {
		if strings.HasSuffix(m.Name, "Raw") || strings.HasSuffix(m.Name, "WithBody") || m.Type.NumOut() != 2 {
			continue
		}
		if strings.HasPrefix(m.Name, "Read") || strings.HasPrefix(m.Name, "List") || strings.HasPrefix(m.Name, "Get") {
			b.BuildMethod(build, m)
		}
	}
	for m := range ct.Methods() {
		if strings.HasSuffix(m.Name, "Raw") || strings.HasSuffix(m.Name, "WithBody") || m.Type.NumOut() != 2 {
			continue
		}
		for _, prefix := range []string{"Delete", "Update", "Put", "Create"} {
			if strings.HasPrefix(m.Name, prefix) {
				b.BuildMethod(build, m)
			}
		}
	}
}

func (b *ClientBuilder[T]) BuildMethod(build *config.Config, m reflect.Method) {
	mb := NewMethodBuilder(b.cfg, build, m)
	err := mb.Build()
	switch {
	case errors.Is(err, ErrCantBuild):
	case err != nil:
		messages.ExitErr(err)
	default:
		mb.Commit()
	}
}
