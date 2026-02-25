/*
SPDX-FileCopyrightText: 2025 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package read

import (
	"reflect"
)

type OOSPager struct{}

func (p OOSPager) HasMore(res reflect.Value) bool {
	if res.Kind() != reflect.Struct {
		return false
	}
	isTruncated := res.FieldByName("IsTruncated")
	if !isTruncated.IsValid() || isTruncated.IsNil() {
		return false
	}
	isTruncated = reflect.Indirect(isTruncated)
	return isTruncated.Kind() == reflect.Bool && isTruncated.Bool()
}

func (OOSPager) nextToken(res reflect.Value) string {
	if res.Kind() != reflect.Struct {
		return ""
	}
	nextToken := res.FieldByName("NextContinuationToken")
	if !nextToken.IsValid() || nextToken.IsNil() {
		return ""
	}
	nextToken = reflect.Indirect(nextToken)
	if nextToken.Kind() != reflect.String {
		return ""
	}
	return nextToken.String()
}

func (p OOSPager) NextItem(res reflect.Value, fetch FetchPage, _ int) (FetchPage, bool) {
	nextToken := p.nextToken(res)
	if nextToken == "" {
		return fetch, false
	}
	for _, arg := range fetch.Args {
		arg = reflect.Indirect(arg)
		if arg.Kind() != reflect.Struct {
			continue
		}
		token := arg.FieldByName("ContinuationToken")
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

var _ Pager = OOSPager{}
