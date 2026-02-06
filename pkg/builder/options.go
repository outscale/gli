/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package builder

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/outscale/gli/pkg/debug"
	"github.com/samber/lo"
)

var NumEntriesInSlices = 1

// We parse the command line to find index-based flags and set NumEntriesInSlices accordingly.
// The cobra commands will be build with all the necessary flags (+1 to allow autompletion of next)
func init() {
	numEntries := 0
	// count the number of flags
	cnt := lo.CountBy(os.Args, func(arg string) bool {
		return strings.HasPrefix(arg, "-")
	})
	// worst case = 1 index per flag
	for i := range cnt {
		idxStr := "." + strconv.Itoa(i) + "."
		if !slices.ContainsFunc(os.Args, func(arg string) bool {
			return strings.Contains(arg, idxStr)
		}) {
			break
		}
		numEntries = i + 1
	}
	NumEntriesInSlices = numEntries + 1
	debug.Println("NumEntriesInSlices", NumEntriesInSlices)
}
