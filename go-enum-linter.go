package main

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var EnumRestrictionAnalyzer = &analysis.Analyzer{
	Name: "enumrestriction",
	Doc:  "restrict enum usage outside specific files",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	allowedFileSuffix := ".enum.go"
	restrictedTypes := map[string]bool{}

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}
	inspector.Preorder(nodeFilter, func(n ast.Node) {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			return
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			if strings.HasSuffix(pass.Fset.Position(n.Pos()).Filename, allowedFileSuffix) {
				restrictedTypes[typeSpec.Name.Name] = true
			}
		}
	})

	inspector.Preorder([]ast.Node{(*ast.CallExpr)(nil), (*ast.AssignStmt)(nil)}, func(n ast.Node) {
		file := pass.Fset.Position(n.Pos()).Filename

		if strings.HasSuffix(file, allowedFileSuffix) {
			return // allows definition file (*.enum.go by default)
		}

		switch stmt := n.(type) {
		case *ast.AssignStmt:
			for _, expr := range stmt.Rhs {
				checkRestrictedType(pass, expr, restrictedTypes)
			}
		case *ast.CallExpr:
			checkRestrictedType(pass, stmt, restrictedTypes)
		}
	})

	return nil, nil
}

func checkRestrictedType(pass *analysis.Pass, expr ast.Expr, restrictedTypes map[string]bool) {
	typ := pass.TypesInfo.TypeOf(expr)
	if named, ok := typ.(*types.Named); ok && restrictedTypes[named.Obj().Name()] {
		pass.Reportf(expr.Pos(), "restricted type %s cannot be used outside .enum.go files", named.Obj().Name())
	}
}
