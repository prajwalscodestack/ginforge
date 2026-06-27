package architecture

type Architecture interface {
	Name() string

	Directories() []string

	ProjectTemplates() []TemplateFile

	GenerateModule(
		projectPath,
		moduleName string,
	) error
}
