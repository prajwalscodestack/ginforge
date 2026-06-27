package architecture

type Architecture interface {
	Name() string

	Directories() []string

	ProjectTemplates() []TemplateFile

	ModuleTemplates(moduleName string) []TemplateFile
}
