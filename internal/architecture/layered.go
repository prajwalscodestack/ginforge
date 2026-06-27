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
func (l Layered) ModuleTemplates(
	moduleName string,
) []TemplateFile {

	return []TemplateFile{
		{
			Template: "templates/layered/module/handler.go.tmpl",
			Output:   "internal/handler/" + moduleName + ".go",
		},
		{
			Template: "templates/layered/module/service.go.tmpl",
			Output:   "internal/service/" + moduleName + ".go",
		},
		{
			Template: "templates/layered/module/repository.go.tmpl",
			Output:   "internal/repository/" + moduleName + ".go",
		},
		{
			Template: "templates/layered/module/model.go.tmpl",
			Output:   "internal/model/" + moduleName + ".go",
		},
		{
			Template: "templates/layered/module/routes.go.tmpl",
			Output:   "internal/routes/" + moduleName + ".go",
		},
	}
}
