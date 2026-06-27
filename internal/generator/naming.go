package generator

import "strings"

func ToPascalCase(s string) string {

	parts := strings.Split(s, "_")

	for i := range parts {

		if len(parts[i]) == 0 {
			continue
		}

		parts[i] =
			strings.ToUpper(parts[i][:1]) +
				parts[i][1:]
	}

	return strings.Join(parts, "")
}
