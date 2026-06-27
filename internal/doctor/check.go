package doctor

type Check interface {
	Name() string
	Run(projectPath string) Result
}
