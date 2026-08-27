package filter

import (
	"context"
	"crypto/sha1" //nolint: gosec
	"encoding/json"
	"iter"
	"slices"

	"github.com/outscale/octl/pkg/output/result"
)

type Dedup struct {
	iter    int
	isNew   bool
	last    [][sha1.Size]byte
	current [][sha1.Size]byte
}

func (j *Dedup) Filter(ctx context.Context, seq iter.Seq[result.Result]) iter.Seq[result.Result] {
	return func(yield func(result.Result) bool) {
		for v := range seq {
			if v.Error != nil {
				_ = yield(v)
				return
			}
			if v.Iter != j.iter {
				j.iter = v.Iter
				j.isNew = false
				j.current, j.last = j.last, j.current
				j.current = j.current[:0]
			}
			js, err := json.Marshal(v.Ok)
			if err == nil {
				hash := sha1.Sum(js) //nolint: gosec
				j.current = append(j.current, hash)
				if !j.isNew && slices.ContainsFunc(j.last, func(e [sha1.Size]byte) bool {
					return hash == e
				}) {
					continue
				}
				j.isNew = true
			}
			if !yield(v) {
				return
			}
		}
	}
}

const dedupBuffer = 100

func NewDedup(resetOnNew bool) *Dedup {
	return &Dedup{last: make([][20]byte, dedupBuffer), current: make([][20]byte, dedupBuffer)}
}

var _ Interface = (*Dedup)(nil)
