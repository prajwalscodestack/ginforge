package architecture

import (
	"github.com/prajwalscodestack/ginforge/internal/config"
)

func ResolveFromProject(
	projectPath string,
) (Architecture, error) {

	cfg, err := config.Load(
		projectPath + "/.ginforge.yaml",
	)
	if err != nil {
		return nil, err
	}

	return Get(cfg.Architecture)
}
