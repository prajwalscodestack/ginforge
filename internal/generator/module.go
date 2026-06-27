package generator

import (
	"path/filepath"

	"ginforge/internal/architecture"
)

func GenerateModule(
	projectPath string,
	moduleName string,
	arch architecture.Architecture,
) error {

	data := ModuleTemplateData{
		ModuleName: moduleName,
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
