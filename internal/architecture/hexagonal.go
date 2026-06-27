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
func (h Hexagonal) ModuleTemplates(
	moduleName string,
) []TemplateFile {

	return []TemplateFile{
		{
			Template: "templates/hexagonal/module/domain.go.tmpl",
			Output: "internal/domain/" +
				moduleName +
				".go",
		},
		{
			Template: "templates/hexagonal/module/application.go.tmpl",
			Output: "internal/application/" +
				moduleName +
				"_service.go",
		},
		{
			Template: "templates/hexagonal/module/inbound.go.tmpl",
			Output: "internal/ports/inbound/" +
				moduleName +
				"_service.go",
		},
		{
			Template: "templates/hexagonal/module/outbound.go.tmpl",
			Output: "internal/ports/outbound/" +
				moduleName +
				"_repository.go",
		},
		{
			Template: "templates/hexagonal/module/handler.go.tmpl",
			Output: "internal/adapters/http/" +
				moduleName +
				"_handler.go",
		},
		{
			Template: "templates/hexagonal/module/repository.go.tmpl",
			Output: "internal/adapters/repository/" +
				moduleName +
				"_repository.go",
		},
	}
}
