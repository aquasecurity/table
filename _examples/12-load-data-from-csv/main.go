package main

import (
	"os"

	"github.com/aquasecurity/table"
)

func main() {

	f, err := os.Open("./_examples/12-load-data-from-csv/data.csv")
	if err != nil {
		panic(err)
	}

	t := table.New(os.Stdout)
	if err := t.LoadCSV(f, true); err != nil {
		panic(err)
	}
	t.Render()
}
