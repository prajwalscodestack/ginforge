package doctor

import (
	"os"
	"path/filepath"
	"strings"
)

func detectArchitecture(
	projectPath string,
) string {

	configPath := filepath.Join(
		projectPath,
		".ginforge.yaml",
	)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return ""
	}

	content := string(data)

	switch {

	case strings.Contains(
		content,
		"architecture: layered",
	):
		return "layered"

	case strings.Contains(
		content,
		"architecture: hexagonal",
	):
		return "hexagonal"
	}

	return ""
}

func getRequiredDirectories(
	architecture string,
) []string {

	switch architecture {

	case "layered":

		return []string{
			"internal/handler",
			"internal/service",
			"internal/repository",
			"internal/model",
			"internal/routes",
		}

	case "hexagonal":

		return []string{
			"internal/domain",
			"internal/application",
			"internal/adapters",
		}
	}

	return nil
}
