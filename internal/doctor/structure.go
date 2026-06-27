package doctor

import (
	"fmt"
	"os"
	"path/filepath"
)

type StructureCheck struct{}

func (StructureCheck) Name() string {
	return "Structure"
}

func (StructureCheck) Run(
	projectPath string,
) Result {

	architecture := detectArchitecture(projectPath)

	if architecture == "" {

		return Result{
			Name:     "Structure",
			Passed:   true,
			Messages: []string{"Required directories present"},
		}
	}

	requiredDirs := getRequiredDirectories(
		architecture,
	)

	var missing []string

	for _, dir := range requiredDirs {

		fullPath := filepath.Join(
			projectPath,
			dir,
		)

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			missing = append(
				missing,
				dir,
			)
		}
	}

	if len(missing) > 0 {

		return Result{
			Name:   "Structure",
			Passed: false,
			Messages: []string{fmt.Sprintf(
				"Missing directories: %v",
				missing,
			)},
		}
	}

	return Result{
		Name:     "Structure",
		Passed:   true,
		Messages: []string{"Required directories present"},
	}
}
