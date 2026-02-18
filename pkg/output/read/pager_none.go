/*
SPDX-FileCopyrightText: 2025 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package read

import "reflect"

type NonePager struct{}

func (p NonePager) HasMore(res reflect.Value) bool {
	return false
}

func (p NonePager) NextItem(res reflect.Value, fetch FetchPage, _ int) (FetchPage, bool) {
	return fetch, false
}

var _ Pager = NonePager{}
