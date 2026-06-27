package generator

import (
	"path/filepath"

	"github.com/prajwalscodestack/ginforge/internal/architecture"
	"github.com/prajwalscodestack/ginforge/internal/config"
)

func GenerateModule(
	projectPath string,
	moduleName string,
	arch architecture.Architecture,
) error {

	modulePath, err := config.ModulePath(projectPath)
	if err != nil {
		return err
	}

	data := ModuleTemplateData{
		ModuleName: moduleName,
		EntityName: ToPascalCase(moduleName),
		ModulePath: modulePath,
	}
	for _, file := range arch.ModuleTemplates(moduleName) {

		if err := RenderTemplate(
			file.Template,
			filepath.Join(projectPath, file.Output),
			data,
		); err != nil {
			return err
		}
	}

	return nil
}
