package config

import (
	"bufio"
	"os"
	"strings"
)

func ModulePath(projectPath string) (string, error) {

	file, err := os.Open(projectPath + "/go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "module ") {

			return strings.TrimSpace(
				strings.TrimPrefix(
					line,
					"module ",
				),
			), nil
		}
	}

	return "", scanner.Err()
}
