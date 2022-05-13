package main

import (
	"os"
	"time"

	"github.com/aquasecurity/table"
)

func main() {

	t := table.New(os.Stdout)

	t.SetAutoMerge(true)

	t.SetHeaders("System", "Status", "Last Check")

	t.AddRow("Life Support", "OK", time.Now().Format(time.Stamp))
	t.AddRow("Nuclear Generator", "OK", time.Now().Add(-time.Minute).Format(time.Stamp))
	t.AddRow("Weapons Systems", "FAIL", time.Now().Format(time.Stamp))
	t.AddRow("Shields", "OK", time.Now().Format(time.Stamp))

	t.Render()
}
