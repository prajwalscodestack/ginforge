package doctor

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type DependencyCheck struct{}

func (DependencyCheck) Name() string {
	return "Dependency Rules"
}

func (DependencyCheck) Run(
	projectPath string,
) Result {

	architecture :=
		detectArchitecture(projectPath)

	if architecture == "" {

		return Result{
			Name:   "Dependency Rules",
			Passed: true,
			Messages: []string{
				"Architecture unknown, dependency validation skipped",
			},
		}
	}

	var violations []string

	err := filepath.Walk(
		filepath.Join(projectPath, "internal"),
		func(
			path string,
			info os.FileInfo,
			err error,
		) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if !strings.HasSuffix(path, ".go") {
				return nil
			}

			fileViolations, err :=
				validateFileDependencies(
					path,
					architecture,
				)

			if err != nil {
				return err
			}

			violations = append(
				violations,
				fileViolations...,
			)

			return nil
		},
	)

	if err != nil {

		return Result{
			Name:   "Dependency Rules",
			Passed: false,
			Messages: []string{
				err.Error(),
			},
		}
	}

	if len(violations) > 0 {

		return Result{
			Name:     "Dependency Rules",
			Passed:   false,
			Messages: violations,
		}
	}

	return Result{
		Name:   "Dependency Rules",
		Passed: true,
		Messages: []string{
			"No dependency violations found",
		},
	}
}

func validateFileDependencies(
	filePath string,
	architecture string,
) ([]string, error) {

	var violations []string

	fset := token.NewFileSet()

	file, err := parser.ParseFile(
		fset,
		filePath,
		nil,
		parser.ImportsOnly,
	)
	if err != nil {
		return nil, err
	}

	switch architecture {

	case "layered":

		layer := detectLayer(filePath)

		for _, imp := range file.Imports {

			importPath :=
				strings.Trim(
					imp.Path.Value,
					"\"",
				)

			importedLayer :=
				extractLayer(importPath)

			if importedLayer == "" {
				continue
			}

			if isForbiddenImport(
				layer,
				importedLayer,
			) {

				violations = append(
					violations,
					fmt.Sprintf(
						"%s\n    %s layer must not import %s layer",
						filePath,
						layer,
						importedLayer,
					),
				)
			}
		}

	case "hexagonal":

		for _, imp := range file.Imports {

			importPath :=
				strings.Trim(
					imp.Path.Value,
					"\"",
				)

			if strings.Contains(
				filePath,
				"/domain/",
			) {

				if strings.Contains(
					importPath,
					"github.com/gin-gonic/gin",
				) {

					violations = append(
						violations,
						fmt.Sprintf(
							"%s\n    domain layer must not depend on gin",
							filePath,
						),
					)
				}

				if strings.Contains(
					importPath,
					"database/sql",
				) {

					violations = append(
						violations,
						fmt.Sprintf(
							"%s\n    domain layer must not depend on database/sql",
							filePath,
						),
					)
				}
			}

			if strings.Contains(
				filePath,
				"/application/",
			) {

				if strings.Contains(
					importPath,
					"/adapters",
				) {

					violations = append(
						violations,
						fmt.Sprintf(
							"%s\n    application layer must not import adapters",
							filePath,
						),
					)
				}
			}
		}
	}

	return violations, nil
}

func detectLayer(path string) string {

	switch {

	case strings.Contains(path, "/handler/"):
		return "handler"

	case strings.Contains(path, "/service/"):
		return "service"

	case strings.Contains(path, "/repository/"):
		return "repository"

	case strings.Contains(path, "/model/"):
		return "model"
	}

	return ""
}

func extractLayer(
	importPath string,
) string {

	switch {

	case strings.Contains(
		importPath,
		"/handler",
	):
		return "handler"

	case strings.Contains(
		importPath,
		"/service",
	):
		return "service"

	case strings.Contains(
		importPath,
		"/repository",
	):
		return "repository"

	case strings.Contains(
		importPath,
		"/model",
	):
		return "model"
	}

	return ""
}

func isForbiddenImport(
	layer string,
	importedLayer string,
) bool {

	rules := map[string][]string{
		"service": {
			"handler",
		},
		"repository": {
			"handler",
			"service",
		},
		"model": {
			"handler",
			"service",
			"repository",
		},
	}

	for _, forbidden := range rules[layer] {

		if forbidden == importedLayer {
			return true
		}
	}

	return false
}
