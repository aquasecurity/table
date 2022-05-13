package main

// noreadme

import (
	"os"

	"github.com/aquasecurity/table"
)

func main() {
	t := table.New(os.Stdout)
	t.SetHeaders("A", "B", "C")
	t.AddRow("🔥 unicode 🔥 characters 🔥", "2", "3")
	t.AddRow("4", "5", "6")
	t.Render()
}
