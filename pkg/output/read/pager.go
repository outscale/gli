/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package read

import (
	"reflect"

	"github.com/outscale/octl/pkg/debug"
)

type Pager interface {
	HasMore(res reflect.Value) bool
	NextItem(res reflect.Value, fetch FetchPage, nextIndex int) (FetchPage, bool)
}

func PagerFor(fetch FetchPage) Pager {
	for _, arg := range fetch.Args {
		arg = reflect.Indirect(arg)
		if arg.Kind() != reflect.Struct {
			continue
		}
		firstItem := arg.FieldByName("FirstItem")
		if firstItem.IsValid() {
			debug.Println("paging using FirstItem")
			return FirstItemPager{}
		}
		nextToken := arg.FieldByName("NextPageToken")
		if nextToken.IsValid() {
			debug.Println("paging using NextPageToken")
			return TokenPager{}
		}
	}
	debug.Println("no paging")
	return NonePager{}
}
