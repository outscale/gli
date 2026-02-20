/*
SPDX-FileCopyrightText: 2025 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package read

import (
	"context"
	"errors"
	"iter"
	"reflect"

	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/output/result"
)

const MaxPages = 20

type Paginated struct {
	contentField string
}

func NewPaginated(contentField string) *Paginated {
	return &Paginated{
		contentField: contentField,
	}
}

func (p *Paginated) Read(ctx context.Context, fetch FetchPage) iter.Seq[result.Result] {
	pager := PagerFor(fetch)
	return func(yield func(result.Result) bool) {
		fetch := fetch
		idx := 0
		for range MaxPages {
			vres := fetch.Call(ctx)
			if len(vres) == 0 {
				_ = yield(result.Result{Error: errors.New("no result from call")})
				return
			}
			if err, ok := vres[len(vres)-1].Interface().(error); ok && err != nil {
				_ = yield(result.Result{Error: err})
				return
			}
			if len(vres) < 2 {
				debug.Println("not enough result")
				_ = yield(result.Result{SingleEntry: true})
				return
			}
			res := reflect.Indirect(vres[0])
			content := res
			if p.contentField != "" {
				content = reflect.Indirect(content.FieldByName(p.contentField))
			}
			if content.Kind() != reflect.Slice {
				debug.Println("not a slice")
				_ = yield(result.Result{Ok: content.Interface(), SingleEntry: true})
				return
			}
			for i := range content.Len() {
				if !yield(result.Result{Ok: content.Index(i).Interface()}) {
					debug.Println("end")
					return
				}
			}
			idx += content.Len()
			if !pager.HasMore(res) {
				debug.Println("end of list...")
				return
			}
			debug.Println("has more items...")
			var ok bool
			fetch, ok = pager.NextItem(res, fetch, idx)
			if !ok {
				return
			}
		}
	}
}

var _ Interface = (*Paginated)(nil)
