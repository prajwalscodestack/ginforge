package architecture

import (
	"os"
	"path/filepath"
)

func CreateDirectories(root string, dirs []string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(
			filepath.Join(root, dir),
			0755,
		); err != nil {
			return err
		}
	}

	return nil
}
