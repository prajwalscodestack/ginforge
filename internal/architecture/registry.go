package architecture

import "fmt"

func Get(name string) (Architecture, error) {
	switch name {

	case "layered":
		return Layered{}, nil

	case "hexagonal":
		return Hexagonal{}, nil

	default:
		return nil, fmt.Errorf("unsupported architecture: %s", name)
	}
}
