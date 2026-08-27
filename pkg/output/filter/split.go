package filter

import (
	"bufio"
	"context"
	"fmt"
	"iter"
	"strings"

	"github.com/outscale/octl/pkg/output/result"
)

type Split struct{}

func (j Split) Filter(ctx context.Context, seq iter.Seq[result.Result]) iter.Seq[result.Result] {
	return func(yield func(result.Result) bool) {
		for v := range seq {
			if v.Error != nil {
				_ = yield(v)
				return
			}
			str, ok := v.Ok.(string)

			if !ok {
				if !yield(v) {
					return
				}
				continue
			}

			scanner := bufio.NewScanner(strings.NewReader(str))
			for scanner.Scan() {
				if !yield(result.New(v, scanner.Text())) {
					return
				}
			}
			if err := scanner.Err(); err != nil {
				_ = yield(result.Error(fmt.Errorf("split: %w", err)))
				return
			}
		}
	}
}

func NewSplit() Split {
	return Split{}
}

var _ Interface = Split{}
