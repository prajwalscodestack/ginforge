package architecture

type Layered struct{}

func (l Layered) Name() string {
	return "layered"
}

func (l Layered) Directories() []string {
	return []string{
		"cmd/server",

		"internal/handler",
		"internal/service",
		"internal/repository",
		"internal/model",
		"internal/routes",

		"configs",
	}
}

func (l Layered) ProjectTemplates() []TemplateFile {
	return []TemplateFile{
		{
			Template: "templates/layered/main.go.tmpl",
			Output:   "cmd/server/main.go",
		},
	}
}
func (l Layered) GenerateModule(projectPath, moduleName string) error {
	return nil
}
