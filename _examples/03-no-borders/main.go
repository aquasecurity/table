package main

import (
	"os"

	"github.com/aquasecurity/table"
)

func main() {

	t := table.New(os.Stdout)
	t.SetBorders(false)

	t.SetHeaders("ID", "Fruit", "Stock")

	t.AddRow("1", "Apple", "14")
	t.AddRow("2", "Banana", "88,041")
	t.AddRow("3", "Cherry", "342")
	t.AddRow("4", "Dragonfruit", "1")

	t.Render()
}
