package doctor

func Run(projectPath string) []Result {

	checks := []Check{
		ArchitectureCheck{},
		StructureCheck{},
		DuplicateRouteCheck{},
	}

	var results []Result

	for _, check := range checks {
		results = append(
			results,
			check.Run(projectPath),
		)
	}

	return results
}
