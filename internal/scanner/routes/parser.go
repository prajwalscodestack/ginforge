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

	ast.Inspect(file, func(n ast.Node) bool {

		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
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

		arg, ok := callExpr.Args[0].(*ast.BasicLit)
		if !ok {
			return true
		}

		pathValue := strings.Trim(arg.Value, "\"")

		routes = append(routes, Route{
			Method: method,
			Path:   pathValue,
			File:   path,
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
