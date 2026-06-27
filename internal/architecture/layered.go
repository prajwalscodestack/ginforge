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

func (l Layered) GenerateProject(path string) error {
	return CreateDirectories(path, l.Directories())
}

func (l Layered) GenerateModule(projectPath, moduleName string) error {
	return nil
}
