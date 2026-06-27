package architecture

type Architecture interface {
	Name() string

	Directories() []string

	GenerateProject(path string) error

	GenerateModule(projectPath, moduleName string) error
}
