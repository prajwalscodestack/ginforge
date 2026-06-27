package routes

import (
	"io/fs"
	"path/filepath"
)

func Scan(projectPath string) ([]Route, error) {

	var routes []Route

	err := filepath.Walk(
		projectPath,
		func(path string, info fs.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if filepath.Ext(path) != ".go" {
				return nil
			}

			fileRoutes, err := ParseFile(path)
			if err != nil {
				return err
			}

			routes = append(routes, fileRoutes...)

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return routes, nil
}
