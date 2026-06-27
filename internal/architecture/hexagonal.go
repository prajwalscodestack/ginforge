package architecture

type Hexagonal struct{}

func (h Hexagonal) Name() string {
	return "hexagonal"
}

func (h Hexagonal) Directories() []string {
	return []string{
		"cmd/server",

		"internal/domain",

		"internal/ports/inbound",
		"internal/ports/outbound",

		"internal/application",

		"internal/adapters/http",
		"internal/adapters/repository",
		"internal/adapters/database",

		"configs",
		"pkg",
	}
}
func (h Hexagonal) ProjectTemplates() []TemplateFile {
	return []TemplateFile{
		{
			Template: "templates/hexagonal/main.go.tmpl",
			Output:   "cmd/server/main.go",
		},
	}
}
func (h Hexagonal) GenerateModule(projectPath, moduleName string) error {
	return nil
}
