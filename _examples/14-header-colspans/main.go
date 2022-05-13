package main

import (
	"os"

	"github.com/aquasecurity/table"
)

func main() {
	t := table.New(os.Stdout)
	t.SetHeaders("Namespace", "Resource", "Vulnerabilities", "Misconfigurations")
	t.AddHeaders("Namespace", "Resource", "Critical", "High", "Medium", "Low", "Unknown", "Critical", "High", "Medium", "Low", "Unknown")
	t.SetHeaderColSpans(0, 1, 1, 5, 5)
	t.SetAutoMergeHeaders(true)
	t.AddRow("default", "Deployment/app", "2", "5", "7", "8", "0", "0", "3", "5", "19", "0")
	t.AddRow("default", "Ingress/test", "-", "-", "-", "-", "-", "1", "0", "2", "17", "0")
	t.AddRow("default", "Service/test", "0", "0", "0", "1", "0", "3", "0", "4", "9", "0")
	t.Render()
}
