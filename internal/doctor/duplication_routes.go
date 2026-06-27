package doctor

import (
	"fmt"

	routescanner "github.com/prajwalscodestack/ginforge/internal/scanner/routes"
)

type DuplicateRouteCheck struct{}

func (DuplicateRouteCheck) Name() string {
	return "Duplicate Routes"
}

func (DuplicateRouteCheck) Run(
	projectPath string,
) Result {

	routes, err :=
		routescanner.Scan(projectPath)

	if err != nil {

		return Result{
			Name:     "Duplicate Routes",
			Passed:   false,
			Messages: []string{err.Error()},
		}
	}

	routeMap :=
		make(map[string]int)

	for _, route := range routes {

		key := fmt.Sprintf(
			"%s:%s",
			route.Method,
			route.Path,
		)

		routeMap[key]++
	}

	for key, count := range routeMap {

		if count > 1 {

			return Result{
				Name:   "Duplicate Routes",
				Passed: false,
				Messages: []string{fmt.Sprintf(
					"Duplicate route found (%s)",
					key,
				)},
			}
		}
	}

	return Result{
		Name:     "Duplicate Routes",
		Passed:   true,
		Messages: []string{"No duplicate routes found"},
	}
}
