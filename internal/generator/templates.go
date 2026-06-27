package generator

import (
	"embed"
	"os"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

type ProjectTemplateData struct {
	ProjectName  string
	Architecture string
	ModuleName   string
}
type ModuleTemplateData struct {
	ModuleName string
	EntityName string
	ModulePath string
}

func RenderTemplate(
	templatePath string,
	outputPath string,
	data any,
) error {

	tmpl, err := template.ParseFS(
		templateFS,
		templatePath,
	)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
