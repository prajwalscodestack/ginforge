package generator

import (
	"path/filepath"

	"ginforge/internal/architecture"
)

func GenerateProject(
	projectPath string,
	arch architecture.Architecture,
) error {

	if err := arch.GenerateProject(projectPath); err != nil {
		return err
	}

	data := ProjectTemplateData{
		ProjectName:  filepath.Base(projectPath),
		Architecture: arch.Name(),
	}

	files := []struct {
		Template string
		Output   string
	}{
		{
			"templates/main.go.tmpl",
			filepath.Join(
				projectPath,
				"cmd/server/main.go",
			),
		},
		{
			"templates/readme.md.tmpl",
			filepath.Join(
				projectPath,
				"README.md",
			),
		},
		{
			"templates/gitignore.tmpl",
			filepath.Join(
				projectPath,
				".gitignore",
			),
		},
		{
			"templates/config.yaml.tmpl",
			filepath.Join(
				projectPath,
				"configs/config.yaml",
			),
		},
	}

	for _, f := range files {
		if err := RenderTemplate(
			f.Template,
			f.Output,
			data,
		); err != nil {
			return err
		}
	}

	return nil
}
