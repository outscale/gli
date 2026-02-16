/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package openapi

import (
	"bufio"
	"regexp"
	"slices"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Spec struct {
	spec *openapi3.T
}

func NewSpec(spec *openapi3.T) *Spec {
	return &Spec{spec: spec}
}

func (s *Spec) SummaryForOperation(name string) (short, help, group string, deprecated bool) {
	if s.spec == nil {
		return "", "", "", false
	}
	for _, p := range s.spec.Paths.Map() {
		for _, op := range p.Operations() {
			if op.OperationID == name {
				help = clean(op.Description)
				r := bufio.NewScanner(strings.NewReader(help))
				for r.Scan() {
					if strings.HasPrefix(r.Text(), "> ") || strings.TrimSpace(r.Text()) == "" {
						continue
					}
					short = r.Text()
					break
				}
				if len(op.Tags) > 0 {
					group = op.Tags[0]
				}
				deprecated = op.Deprecated
				return
			}
		}
	}
	return "", "", "", false
}

func (s *Spec) SummaryForAttribute(typeName, attribute string) (description string, required bool) {
	if s.spec == nil {
		return "", false
	}
	typeDef := s.spec.Components.Schemas[typeName]
	if typeDef == nil || typeDef.Value == nil {
		return "", false
	}

	propDef := typeDef.Value.Properties[attribute]
	if propDef == nil || propDef.Value == nil {
		return "", false
	}
	return clean(propDef.Value.Description), slices.Contains(typeDef.Value.Required, attribute)
}

var reEOL = regexp.MustCompile("[\r\n]{2,}")

func clean(str string) string {
	r := strings.NewReplacer(
		"<br />", "\n",
		"\\|", "|",
		"`", "",
		"\r\n", "\n",
	)
	return reEOL.ReplaceAllString(r.Replace(str), "\n\n")
}
