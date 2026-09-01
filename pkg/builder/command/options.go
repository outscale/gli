/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package commandbuilder

import (
	"os"
	"strconv"
	"strings"

	"github.com/outscale/octl/pkg/debug"
	"github.com/samber/lo"
)

var numEntriesInSlices = map[string]int{}

// We parse the command line to find index-based flags and set NumEntriesInSlices accordingly.
// The cobra commands will be build with all the necessary flags (+1 to allow autompletion of next)
func init() {
	// count the number of flags
	cnt := lo.CountBy(os.Args, func(arg string) bool {
		return strings.HasPrefix(arg, "--")
	})
	// worst case = 1 index per flag
	for i := range cnt {
		idxStr := "." + strconv.Itoa(i) + "."
		for _, arg := range os.Args {
			parts := strings.Split(strings.TrimPrefix(arg, "--"), idxStr)
			if len(parts) == 1 {
				continue
			}
			prefix := ""
			for iarg := range len(parts) - 1 {
				numEntriesInSlices[prefix+parts[iarg]] = i + 1
				prefix += parts[iarg] + idxStr
			}
		}
	}
	debug.Println("NumEntriesInSlices", numEntriesInSlices)
}

func NumEntriesInSlices(prefix string) int {
	if n, found := numEntriesInSlices[prefix]; found {
		return n + 1
	}
	return 1
}
