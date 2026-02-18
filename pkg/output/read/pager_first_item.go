/*
SPDX-FileCopyrightText: 2025 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package read

import "reflect"

type FirstItemPager struct{}

func (FirstItemPager) HasMore(res reflect.Value) bool {
	if res.Kind() != reflect.Struct {
		return false
	}
	hasMore := res.FieldByName("HasMoreItems")
	if !hasMore.IsValid() || hasMore.IsNil() {
		return false
	}
	hasMore = reflect.Indirect(hasMore)
	return hasMore.Kind() == reflect.Bool && hasMore.Bool()
}

func (FirstItemPager) NextItem(res reflect.Value, fetch FetchPage, nextIndex int) (FetchPage, bool) {
	for _, arg := range fetch.Args {
		arg = reflect.Indirect(arg)
		if arg.Kind() != reflect.Struct {
			continue
		}
		firstItem := arg.FieldByName("FirstItem")
		if !firstItem.IsValid() {
			continue
		}
		if firstItem.Kind() == reflect.Pointer && firstItem.IsNil() {
			firstItem.Set(reflect.New(firstItem.Type().Elem()))
		}
		firstItem = reflect.Indirect(firstItem)
		if firstItem.Kind() != reflect.Int {
			continue
		}
		firstItem.Set(reflect.ValueOf(nextIndex))
		return fetch, true
	}
	return fetch, false
}

var _ Pager = FirstItemPager{}
