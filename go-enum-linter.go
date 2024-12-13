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

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter = []ast.Node{
		(*ast.ValueSpec)(nil),    // For var and const declarations
		(*ast.CallExpr)(nil),     // For function calls
		(*ast.CompositeLit)(nil), // For struct literals
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.ValueSpec:
			// Check var and const declarations
			for _, name := range node.Names {
				if typ, ok := node.Type.(*ast.Ident); ok {
					if isEnumType(typ.Name) {
						pass.Reportf(name.Pos(), "some")
					}
				}
			}

		case *ast.CallExpr:
			// Check function calls with enum parameters
			if fun, ok := node.Fun.(*ast.Ident); ok {
				if fd := pass.TypesInfo.ObjectOf(fun); fd != nil {
					if fn, ok := fd.(*types.Func); ok {
						sig := fn.Type().(*types.Signature)
						params := sig.Params()
						for i := 0; i < params.Len(); i++ {
							param := params.At(i)
							if isEnumType(param.Type().String()) {
								pass.Reportf(node.Args[i].Pos(), "some")
							}
						}
					}
				}
			}

		case *ast.CompositeLit:
			// Check struct literal field assignments
			if _, ok := node.Type.(*ast.Ident); ok {
				for _, elt := range node.Elts {
					if kv, ok := elt.(*ast.KeyValueExpr); ok {
						if key, ok := kv.Key.(*ast.Ident); ok {
							if isEnumType(key.Name) {
								pass.Reportf(kv.Value.Pos(), "some")
							}
						}
					}
				}
			}
		}
	})

	return nil, nil
}

func isEnumType(name string) bool {
	return name == "Status" || name == "Category"
}

func checkRestrictedType(pass *analysis.Pass, expr ast.Expr, restrictedTypes map[string]bool) {
	typ := pass.TypesInfo.TypeOf(expr)
	if named, ok := typ.(*types.Named); ok && restrictedTypes[named.Obj().Name()] {
		pass.Reportf(expr.Pos(), "restricted type %s cannot be used outside .enum.go files", named.Obj().Name())
	}
}
