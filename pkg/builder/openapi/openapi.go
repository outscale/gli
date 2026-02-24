/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package openapi

import (
	"regexp"
	"slices"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Spec struct {
	spec    *openapi3.T
	helpURL string
}

func NewSpec(spec *openapi3.T, helpURL string) *Spec {
	return &Spec{spec: spec, helpURL: helpURL}
}

func (s *Spec) SummaryForOperation(name string) (short, help, group string, deprecated bool) {
	if s.spec == nil {
		return "", "", "", false
	}
	for _, p := range s.spec.Paths.Map() {
		for _, op := range p.Operations() {
			if op.OperationID == name {
				help = op.Description
				help += "\n\n> Online help: " + s.helpURL + "#" + strings.ToLower(op.OperationID)
				var found bool
				short, _, found = strings.Cut(help, ". ")
				if found {
					short += "."
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
	var found bool
	description, _, found = strings.Cut(cleanOneLine(propDef.Value.Description), ". ")
	if found {
		description += "."
	}
	return description, slices.Contains(typeDef.Value.Required, attribute)
}

var reSpaces = regexp.MustCompile("[ ]{2,}")

func cleanOneLine(str string) string {
	r := strings.NewReplacer(
		"<br />", " ",
		"\\|", "|",
		"`", "",
		"\r", " ",
		"\n", " ",
	)
	return strings.TrimSpace(reSpaces.ReplaceAllString(r.Replace(str), " "))
}
