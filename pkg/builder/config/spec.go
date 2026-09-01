package configbuilder

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/descriptions"
	"github.com/outscale/octl/pkg/messages"
	"github.com/samber/lo"
	"golang.org/x/tools/go/packages"
)

type SpecCall struct {
	Group string
	Help  string
}

type SpecAttribute struct {
	Help     string
	Required bool
}

type Spec struct {
	Calls      map[string]SpecCall      `yaml:"calls,omitempty"`
	Attributes map[string]SpecAttribute `yaml:"attributes,omitempty"`
}

func (s Spec) ForCall(name string) SpecCall {
	if s.Calls == nil {
		return SpecCall{}
	}
	return s.Calls[name]
}

func (s Spec) ForAttribute(call, name string) SpecAttribute {
	if s.Attributes == nil {
		return SpecAttribute{}
	}
	return s.Attributes[call+"."+name]
}

func (s *Spec) SetAttribute(call, name string, spec SpecAttribute) {
	s.Attributes[call+"."+name] = spec
}

type SpecBuilder struct {
	cfg Config
}

func NewSpecBuilder(cfg Config) *SpecBuilder {
	return &SpecBuilder{cfg: cfg}
}

func (b *SpecBuilder) BuildSpec(base *config.Config, pkgNames ...string) Spec {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedSyntax | packages.LoadFiles,
	}, pkgNames...)
	if err != nil {
		messages.ExitErr(err)
	}
	fmt.Println("*** fetching calls from AST")
	spec := Spec{
		Calls:      map[string]SpecCall{},
		Attributes: map[string]SpecAttribute{},
	}

	for _, file := range pkgs[0].Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			if stmt, ok := n.(*ast.FuncDecl); ok && stmt.Recv != nil {
				method := stmt.Name.Name
				if strings.HasSuffix(method, "Raw") || strings.HasSuffix(method, "WithBody") {
					return true
				}
				if !stmt.Name.IsExported() {
					return false
				}
				if len(stmt.Type.Params.List) < 2 {
					return false
				}
				if stmt.Type.Results == nil {
					return false
				}
				if _, found := base.API[method]; !found {
					var group string
					if stmt.Doc != nil {
						groupComment, found := lo.Find(stmt.Doc.List, func(c *ast.Comment) bool {
							return strings.HasPrefix(c.Text, "//sdk:group")
						})
						if found {
							_, group, _ = strings.Cut(groupComment.Text, " ")
						}
					}
					txt := strings.TrimSpace(strings.TrimPrefix(stmt.Doc.Text(), method))
					spec.Calls[method] = SpecCall{
						Group: group,
						Help:  descriptions.Clean(txt),
					}
				}
			}
			return true
		})
	}
	fmt.Println("*** fetching field spec from AST")
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(n ast.Node) bool {
				if stmt, ok := n.(*ast.TypeSpec); ok {
					typeName := stmt.Name.Name
					if strings.HasSuffix(typeName, "Interface") || strings.HasSuffix(typeName, "Raw") ||
						strings.HasSuffix(typeName, "Response") || strings.HasSuffix(typeName, "Output") {
						return false
					}
					if !stmt.Name.IsExported() {
						return false
					}
					ast.Inspect(stmt, func(n ast.Node) bool {
						if stmt, ok := n.(*ast.Field); ok {
							if len(stmt.Names) == 0 { // anonymous field
								return false
							}
							if !stmt.Names[0].IsExported() {
								return false
							}
							flag := stmt.Names[0].Name
							doc := strings.TrimSpace(strings.TrimPrefix(stmt.Doc.Text(), flag))
							help := descriptions.Summary(doc)
							var required bool
							if b.cfg.RequiredFromComment != nil {
								required = b.cfg.RequiredFromComment(doc)
							}
							if help != "" || required {
								spec.SetAttribute(typeName, flag, SpecAttribute{Help: help, Required: required})
							}
							return false
						}
						return true
					})
					return false
				}
				return true
			})
		}
	}
	return spec
}
