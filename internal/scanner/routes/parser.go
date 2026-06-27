package routes

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func ParseFile(path string) ([]Route, error) {

	fset := token.NewFileSet()

	file, err := parser.ParseFile(
		fset,
		path,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		return nil, err
	}

	var routes []Route

	groups := make(map[string]string)

	ast.Inspect(file, func(n ast.Node) bool {

		// Detect groups
		assignStmt, ok := n.(*ast.AssignStmt)

		if ok {

			if len(assignStmt.Lhs) != 1 ||
				len(assignStmt.Rhs) != 1 {
				return true
			}

			varName, ok :=
				assignStmt.Lhs[0].(*ast.Ident)

			if !ok {
				return true
			}

			callExpr, ok :=
				assignStmt.Rhs[0].(*ast.CallExpr)

			if !ok {
				return true
			}

			selectorExpr, ok :=
				callExpr.Fun.(*ast.SelectorExpr)

			if !ok {
				return true
			}

			if selectorExpr.Sel.Name != "Group" {
				return true
			}

			if len(callExpr.Args) == 0 {
				return true
			}

			pathArg, ok :=
				callExpr.Args[0].(*ast.BasicLit)

			if !ok {
				return true
			}

			prefix := strings.Trim(
				pathArg.Value,
				"\"",
			)

			fullPrefix := prefix

			receiver, ok :=
				selectorExpr.X.(*ast.Ident)

			if ok {

				if parentPrefix, exists :=
					groups[receiver.Name]; exists {

					fullPrefix =
						parentPrefix + prefix
				}
			}

			groups[varName.Name] =
				fullPrefix

			return true
		}

		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selectorExpr, ok :=
			callExpr.Fun.(*ast.SelectorExpr)

		if !ok {
			return true
		}

		method := selectorExpr.Sel.Name

		if !isHTTPMethod(method) {
			return true
		}

		if len(callExpr.Args) == 0 {
			return true
		}

		arg, ok :=
			callExpr.Args[0].(*ast.BasicLit)

		if !ok {
			return true
		}

		pathValue := strings.Trim(
			arg.Value,
			"\"",
		)

		handlerName := ""

		if len(callExpr.Args) > 1 {

			switch h := callExpr.Args[1].(type) {

			case *ast.Ident:
				handlerName = h.Name

			case *ast.SelectorExpr:
				handlerName = h.Sel.Name

			case *ast.FuncLit:
				handlerName = "inline-func"

			}
		}

		receiver, ok :=
			selectorExpr.X.(*ast.Ident)

		if ok {

			if prefix, exists :=
				groups[receiver.Name]; exists {

				pathValue =
					prefix + pathValue
			}
		}

		routes = append(routes, Route{
			Method:  method,
			Path:    pathValue,
			File:    path,
			Handler: handlerName,
		})

		return true
	})

	return routes, nil
}

func isHTTPMethod(method string) bool {

	switch method {
	case "GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH":
		return true
	}

	return false
}
