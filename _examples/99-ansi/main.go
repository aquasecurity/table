package main

// noreadme

import (
	"os"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
)

func main() {

	t := table.New(os.Stdout)

	t.SetHeaders("ID", "Fruit", "Stock", "Description")
	t.SetHeaderStyle(table.StyleBold)
	t.SetLineStyle(table.StyleBlue)
	t.SetDividers(table.UnicodeRoundedDividers)

	t.AddRow("1", tml.Sprintf("<green>Apple</green>"), "14", tml.Sprintf("An apple is an edible fruit produced by an apple tree (<italic>Malus domestica</italic>). "))
	t.AddRow("2", tml.Sprintf("<yellow>Banana</yellow>"), "88,041", "A banana is an elongated, edible fruit - botanically a berry.")
	t.AddRow("3", tml.Sprintf("<red>Cherry</red>"), "342", "A cherry is the fruit of many plants of the genus Prunus, and is a fleshy drupe (stone fruit). ")
	t.AddRow("4", tml.Sprintf("<magenta>Dragonfruit</magenta>"), "1", "A dragonfruit is the fruit of several different cactus species indigenous to the Americas.")

	t.Render()
}
