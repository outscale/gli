/*
SPDX-FileCopyrightText: 2025 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package read

import (
	"reflect"
)

type TokenPager struct{}

func (TokenPager) nextToken(res reflect.Value) string {
	if res.Kind() != reflect.Struct {
		return ""
	}
	nextToken := res.FieldByName("NextPageToken")
	if !nextToken.IsValid() || nextToken.IsNil() {
		return ""
	}
	nextToken = reflect.Indirect(nextToken)
	if nextToken.Kind() != reflect.String {
		return ""
	}
	return nextToken.String()
}

func (p TokenPager) HasMore(res reflect.Value) bool {
	return p.nextToken(res) != ""
}

func (p TokenPager) NextItem(res reflect.Value, fetch FetchPage, _ int) (FetchPage, bool) {
	nextToken := p.nextToken(res)
	if nextToken == "" {
		return fetch, false
	}
	for _, arg := range fetch.Args {
		arg = reflect.Indirect(arg)
		if arg.Kind() != reflect.Struct {
			continue
		}
		token := arg.FieldByName("NextPageToken")
		if !token.IsValid() {
			continue
		}
		if token.Kind() == reflect.Pointer && token.IsNil() {
			token.Set(reflect.New(token.Type().Elem()))
		}
		token = reflect.Indirect(token)
		if token.Kind() != reflect.String {
			continue
		}
		token.Set(reflect.ValueOf(nextToken))
		return fetch, true
	}
	return fetch, false
}

var _ Pager = TokenPager{}
