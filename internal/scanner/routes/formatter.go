package routes

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func PrintTable(routes []Route) {

	fmt.Printf("%-8s %s\n", "METHOD", "PATH")

	for _, route := range routes {

		fmt.Printf(
			"%-8s %s\n",
			route.Method,
			route.Path,
		)
	}
}

func PrintJSON(routes []Route) error {

	data, err := json.MarshalIndent(
		routes,
		"",
		"  ",
	)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}

func PrintCSV(routes []Route) error {

	writer := csv.NewWriter(os.Stdout)

	if err := writer.Write(
		[]string{
			"METHOD",
			"PATH",
		},
	); err != nil {
		return err
	}

	for _, route := range routes {

		if err := writer.Write(
			[]string{
				route.Method,
				route.Path,
			},
		); err != nil {
			return err
		}
	}

	writer.Flush()

	return writer.Error()
}

func PrintMarkdown(routes []Route) {

	fmt.Println("| Method | Path |")
	fmt.Println("|--------|------|")

	for _, route := range routes {

		fmt.Printf(
			"| %s | %s |\n",
			route.Method,
			route.Path,
		)
	}
}
