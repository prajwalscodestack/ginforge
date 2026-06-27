package doctor

func Run(projectPath string) []Result {

	checks := []Check{
		ArchitectureCheck{},
		StructureCheck{},
		DuplicateRouteCheck{},
		DependencyCheck{},
	}

	var results []Result

	var hasFailure bool

	for _, check := range checks {

		result := check.Run(projectPath)

		results = append(results, result)

		if !result.Passed {
			hasFailure = true
		}
	}

	// ---------------------------------------
	// OPTIONAL: future CI hook readiness
	// ---------------------------------------
	_ = hasFailure

	return results
}
