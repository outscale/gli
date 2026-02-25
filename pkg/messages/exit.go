/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package messages

import (
	"os"
)

func Exit(code int, format string, a ...any) {
	Err(format, a...)
	os.Exit(code)
}

func ExitErr(err error) {
	Exit(1, "an error occurred: %v", err)
}
