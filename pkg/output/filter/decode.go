package filter

import (
	"context"
	"encoding/base64"
	"fmt"
	"iter"

	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/output/result"
)

type Decode struct{}

func (j Decode) Filter(ctx context.Context, seq iter.Seq[result.Result]) iter.Seq[result.Result] {
	return func(yield func(result.Result) bool) {
		for v := range seq {
			if v.Error != nil {
				_ = yield(v)
				return
			}

			if str, ok := v.Ok.(string); ok {
				buf, err := base64.StdEncoding.DecodeString(str)
				if err != nil {
					_ = yield(result.Error(fmt.Errorf("unable to decode: %w", err)))
					return
				}
				v = result.New(v, string(buf))
			} else {
				debug.Println(fmt.Sprintf("decode: not a string (%T)", v.Ok))
			}

			if !yield(v) {
				return
			}
		}
	}
}

func NewDecode() Decode {
	return Decode{}
}

var _ Interface = Decode{}
