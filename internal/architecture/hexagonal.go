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

func (h Hexagonal) GenerateProject(path string) error {
	return CreateDirectories(path, h.Directories())
}

func (h Hexagonal) GenerateModule(projectPath, moduleName string) error {
	return nil
}
