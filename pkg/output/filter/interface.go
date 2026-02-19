package filter

import (
	"context"
	"iter"

	"github.com/outscale/octl/pkg/output/result"
)

type Interface interface {
	Filter(ctx context.Context, seq iter.Seq[result.Result]) iter.Seq[result.Result]
}
