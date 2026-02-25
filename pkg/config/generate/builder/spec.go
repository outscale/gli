package builder

import (
	"fmt"
	"go/ast"
	"strings"

	"dario.cat/mergo"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/descriptions"
	"github.com/outscale/octl/pkg/messages"
	"github.com/samber/lo"
	"golang.org/x/tools/go/packages"
)

type SpecBuilder struct {
	cfg Config
}

func NewSpecBuilder(cfg Config) *SpecBuilder {
	return &SpecBuilder{cfg: cfg}
}

func (b *SpecBuilder) BuildSpec(base *config.Config, pkgNames ...string) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedSyntax | packages.LoadFiles,
	}, pkgNames...)
	if err != nil {
		messages.ExitErr(err)
	}
	fmt.Println("*** fetching calls from AST")
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
				if _, found := base.Calls[method]; !found {
					_, _, entity := names(method)
					var group string
					if stmt.Doc != nil {
						groupComment, found := lo.Find(stmt.Doc.List, func(c *ast.Comment) bool {
							return strings.HasPrefix(c.Text, "//sdk:group")
						})
						if found {
							_, group, _ = strings.Cut(groupComment.Text, " ")
						}
					}
					call := base.Calls[method]
					newCall := config.Call{
						Entity: entity,
						Group:  group,
					}
					err := mergo.Merge(&call, newCall)
					if err != nil {
						messages.ExitErr(err)
					}
					base.Calls[method] = call
					txt := strings.TrimSpace(strings.TrimPrefix(stmt.Doc.Text(), method))
					base.Spec.Calls[method] = config.SpecCall{
						Help: descriptions.Clean(txt),
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
								base.Spec.SetAttribute(typeName, flag, config.SpecAttribute{Help: help, Required: required})
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
}
