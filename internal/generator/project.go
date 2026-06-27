package generator

import (
	"path/filepath"

	"ginforge/internal/architecture"
	"ginforge/internal/config"
)

func GenerateProject(
	projectPath string,
	arch architecture.Architecture,
) error {

	// create folders
	if err := architecture.CreateDirectories(
		projectPath,
		arch.Directories(),
	); err != nil {
		return err
	}

	data := ProjectTemplateData{
		ProjectName:  filepath.Base(projectPath),
		Architecture: arch.Name(),
		ModuleName:   filepath.Base(projectPath),
	}

	// architecture templates
	for _, file := range arch.ProjectTemplates() {

		if err := RenderTemplate(
			file.Template,
			filepath.Join(projectPath, file.Output),
			data,
		); err != nil {
			return err
		}
	}
	// generate shared templates
	for _, file := range SharedTemplates() {

		if err := RenderTemplate(
			file.Template,
			filepath.Join(projectPath, file.Output),
			data,
		); err != nil {
			return err
		}
	}
	cfg := &config.GinForgeConfig{
		Architecture: arch.Name(),
		Version:      1,
	}

	err := config.Save(
		filepath.Join(
			projectPath,
			".ginforge.yaml",
		),
		cfg,
	)
	if err != nil {
		return err
	}

	return nil
}
func SharedTemplates() []architecture.TemplateFile {
	return []architecture.TemplateFile{
		{
			Template: "templates/shared/readme.md.tmpl",
			Output:   "README.md",
		},
		{
			Template: "templates/shared/gitignore.tmpl",
			Output:   ".gitignore",
		},
		{
			Template: "templates/shared/config.yaml.tmpl",
			Output:   "configs/config.yaml",
		},
		{
			Template: "templates/shared/gomod.tmpl",
			Output:   "go.mod",
		},
	}
}
