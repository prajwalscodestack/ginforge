package doctor

import "fmt"

type ArchitectureCheck struct{}

func (ArchitectureCheck) Name() string {
	return "Architecture"
}
func (ArchitectureCheck) Run(
	projectPath string,
) Result {

	architecture :=
		detectArchitecture(projectPath)

	if architecture == "" {

		return Result{
			Name:     "Architecture",
			Passed:   true,
			Messages: []string{"Standard Gin project detected"},
		}
	}

	return Result{
		Name:   "Architecture",
		Passed: true,
		Messages: []string{fmt.Sprintf(
			"GinForge project detected (%s)",
			architecture,
		)},
	}
}
