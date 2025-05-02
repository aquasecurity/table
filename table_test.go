package table

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertMultilineEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf(`Error: Tables are not equal.

Expected: %s
Actual: %s
`, expected, actual)
	}
}

func Test_BasicTable(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ A │ B │ C │
├───┼───┼───┤
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_EmptyTable(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.Render()
	assertMultilineEqual(t, ``, builder.String())
}

func Test_AddRows(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRows([]string{"1", "2", "3"}, []string{"4", "5", "6"})
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ A │ B │ C │
├───┼───┼───┤
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_NoHeaders(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_Footers(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetFooters("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
├───┼───┼───┤
│ A │ B │ C │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_Footers_No_RowLines(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetFooters("A", "B", "C")
	table.SetRowLines(false)
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.AddRow("7", "8", "9")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ 1 │ 2 │ 3 │
│ 4 │ 5 │ 6 │
│ 7 │ 8 │ 9 │
├───┼───┼───┤
│ A │ B │ C │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_VaryingWidths(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("AAA", "BBBBB", "CCCCCCCCC")
	table.AddRow("111111", "2", "3")
	table.AddRow("4", "5555555555", "6")
	table.Render()
	assertMultilineEqual(t, `
┌────────┬────────────┬───────────┐
│  AAA   │   BBBBB    │ CCCCCCCCC │
├────────┼────────────┼───────────┤
│ 111111 │ 2          │ 3         │
├────────┼────────────┼───────────┤
│ 4      │ 5555555555 │ 6         │
└────────┴────────────┴───────────┘
`, "\n"+builder.String())
}

func Test_Wrapping(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("ID", "Name", "Notes")
	table.AddRow("1", "Jim", "")
	table.AddRow("2", "Bob", "This is a very, very, very, very, very, long sentence which will surely wrap?")
	table.Render()
	assertMultilineEqual(t, `
┌────┬──────┬─────────────────────────────────────────────────────────────┐
│ ID │ Name │                            Notes                            │
├────┼──────┼─────────────────────────────────────────────────────────────┤
│ 1  │ Jim  │                                                             │
├────┼──────┼─────────────────────────────────────────────────────────────┤
│ 2  │ Bob  │ This is a very, very, very, very, very, long sentence which │
│    │      │ will surely wrap?                                           │
└────┴──────┴─────────────────────────────────────────────────────────────┘
`, "\n"+builder.String())
}

func Test_MultipleLines(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("ID", "Name", "Notes")
	table.AddRow("1", "Jim", "")
	table.AddRow("2", "Bob", "This is a sentence.\nThis is another sentence.\nAnd yet another one!")
	table.Render()
	assertMultilineEqual(t, `
┌────┬──────┬───────────────────────────┐
│ ID │ Name │           Notes           │
├────┼──────┼───────────────────────────┤
│ 1  │ Jim  │                           │
├────┼──────┼───────────────────────────┤
│ 2  │ Bob  │ This is a sentence.       │
│    │      │ This is another sentence. │
│    │      │ And yet another one!      │
└────┴──────┴───────────────────────────┘
`, "\n"+builder.String())
}

func Test_Padding_None(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("ID", "Name", "Notes")
	table.AddRow("1", "Jim", "")
	table.AddRow("2", "Bob", "This is a sentence.\nThis is another sentence.\nAnd yet another one!")
	table.SetPadding(0)
	table.Render()
	assertMultilineEqual(t, `
┌──┬────┬─────────────────────────┐
│ID│Name│          Notes          │
├──┼────┼─────────────────────────┤
│1 │Jim │                         │
├──┼────┼─────────────────────────┤
│2 │Bob │This is a sentence.      │
│  │    │This is another sentence.│
│  │    │And yet another one!     │
└──┴────┴─────────────────────────┘
`, "\n"+builder.String())
}

func Test_Padding_Lots(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("ID", "Name", "Notes")
	table.AddRow("1", "Jim", "Good")
	table.AddRow("2", "Bob", "Bad")
	table.SetPadding(10)
	table.Render()
	assertMultilineEqual(t, `
┌──────────────────────┬────────────────────────┬─────────────────────────┐
│          ID          │          Name          │          Notes          │
├──────────────────────┼────────────────────────┼─────────────────────────┤
│          1           │          Jim           │          Good           │
├──────────────────────┼────────────────────────┼─────────────────────────┤
│          2           │          Bob           │          Bad            │
└──────────────────────┴────────────────────────┴─────────────────────────┘
`, "\n"+builder.String())
}

func Test_AlignmentDefaults(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("ID", "Name", "Notes")
	table.AddRow("1", "Jim", "This person has a verrrrry short name.")
	table.AddRow("2", "Bob", "See above.")
	table.AddRow("99999", "John Verylongname", "Bad")
	table.Render()
	assertMultilineEqual(t, `
┌───────┬───────────────────┬────────────────────────────────────────┐
│  ID   │       Name        │                 Notes                  │
├───────┼───────────────────┼────────────────────────────────────────┤
│ 1     │ Jim               │ This person has a verrrrry short name. │
├───────┼───────────────────┼────────────────────────────────────────┤
│ 2     │ Bob               │ See above.                             │
├───────┼───────────────────┼────────────────────────────────────────┤
│ 99999 │ John Verylongname │ Bad                                    │
└───────┴───────────────────┴────────────────────────────────────────┘
`, "\n"+builder.String())
}

func Test_AlignmentCustom(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("ID", "Name", "Notes")
	table.SetFooters("ID", "Name", "Notes")
	table.SetHeaderAlignment(AlignLeft, AlignCenter, AlignRight)
	table.SetAlignment(AlignCenter, AlignRight, AlignLeft)
	table.SetFooterAlignment(AlignRight, AlignLeft, AlignCenter)
	table.AddRow("Please", "be", "aligned")
	table.Render()
	assertMultilineEqual(t, `
┌────────┬──────┬─────────┐
│ ID     │ Name │   Notes │
├────────┼──────┼─────────┤
│ Please │   be │ aligned │
├────────┼──────┼─────────┤
│     ID │ Name │  Notes  │
└────────┴──────┴─────────┘
`, "\n"+builder.String())
}

func Test_NonDefaultDividers(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetDividers(UnicodeRoundedDividers)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
╭───┬───┬───╮
│ A │ B │ C │
├───┼───┼───┤
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
╰───┴───┴───╯
`, "\n"+builder.String())
}

func Test_UnequalRows(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1")
	table.AddRow()
	table.AddRow("", "2", "3")
	table.AddRow("4", "5", "6")
	table.AddRow("7", "8", "9", "10")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┬────┐
│ A │ B │ C │    │
├───┼───┼───┼────┤
│ 1 │   │   │    │
├───┼───┼───┼────┤
│   │   │   │    │
├───┼───┼───┼────┤
│   │ 2 │ 3 │    │
├───┼───┼───┼────┤
│ 4 │ 5 │ 6 │    │
├───┼───┼───┼────┤
│ 7 │ 8 │ 9 │ 10 │
└───┴───┴───┴────┘
`, "\n"+builder.String())
}

func Test_NoBorders(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetBorders(false)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
 A │ B │ C 
───┼───┼───
 1 │ 2 │ 3 
───┼───┼───
 4 │ 5 │ 6 
`, "\n"+builder.String())
}

func Test_NoLeftBorder(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetBorderLeft(false)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
───┬───┬───┐
 A │ B │ C │
───┼───┼───┤
 1 │ 2 │ 3 │
───┼───┼───┤
 4 │ 5 │ 6 │
───┴───┴───┘
`, "\n"+builder.String())
}

func Test_NoRightBorder(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetBorderRight(false)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───
│ A │ B │ C 
├───┼───┼───
│ 1 │ 2 │ 3 
├───┼───┼───
│ 4 │ 5 │ 6 
└───┴───┴───
`, "\n"+builder.String())
}

func Test_NoTopBorder(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetBorderTop(false)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
│ A │ B │ C │
├───┼───┼───┤
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_NoBottomBorder(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetBorderBottom(false)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ A │ B │ C │
├───┼───┼───┤
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
`, "\n"+builder.String())
}

func Test_LineStyle(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetLineStyle(StyleYellow)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assert.Equal(t, "\n"+
		"\x1b[33m┌───┬───┬───┐\x1b[0m\n"+
		"\x1b[33m│\u001B[0m A \u001B[33m│\u001B[0m B \u001B[33m│\u001B[0m C \u001B[33m│\x1b[0m\n"+
		"\x1b[33m├───┼───┼───┤\x1b[0m\n"+
		"\x1b[33m│\u001B[0m 1 \u001B[33m│\u001B[0m 2 \u001B[33m│\u001B[0m 3 \u001B[33m│\x1b[0m\n"+
		"\x1b[33m├───┼───┼───┤\x1b[0m\n"+
		"\x1b[33m│\u001B[0m 4 \u001B[33m│\u001B[0m 5 \u001B[33m│\u001B[0m 6 \u001B[33m│\x1b[0m\n"+
		"\x1b[33m└───┴───┴───┘\x1b[0m\n",
		"\n"+builder.String())
}

func Test_PrerenderedANSI(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "\x1b[37m2\x1b[0m", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ A │ B │ C │
├───┼───┼───┤
│ 1 │ `+"\x1b[37m2\x1b[0m"+` │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_NoRowLines(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetRowLines(false)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.AddRow("7", "8", "9")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ A │ B │ C │
├───┼───┼───┤
│ 1 │ 2 │ 3 │
│ 4 │ 5 │ 6 │
│ 7 │ 8 │ 9 │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_AutoMerge(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetAutoMerge(true)
	table.SetHeaders("A", "B", "3")
	table.AddRow("", "2", "3")
	table.AddRow("", "2", "6")
	table.AddRow("1", "2", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───┬───┐
│ A │ B │ 3 │
├───┼───┼───┤
│   │ 2 │ 3 │
├───┤   ├───┤
│   │   │ 6 │
├───┤   │   │
│ 1 │   │   │
└───┴───┴───┘
`, "\n"+builder.String())
	if t.Failed() {
		fmt.Println(builder.String())
	}
}

func Test_Unicode(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRow("🔥 unicode 🔥 characters 🔥", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()

	/*
		The following may look wrong in your editor,
		but when double-width runes are rendered correctly,
		this is right.
	*/

	assertMultilineEqual(t, `
┌─────────────────────────────┬───┬───┐
│              A              │ B │ C │
├─────────────────────────────┼───┼───┤
│ 🔥 unicode 🔥 characters 🔥 │ 2 │ 3 │
├─────────────────────────────┼───┼───┤
│ 4                           │ 5 │ 6 │
└─────────────────────────────┴───┴───┘
`, "\n"+builder.String())
}

func TestCSV(t *testing.T) {

	input := strings.NewReader(`Id,Date,Message
1,2022-05-12,"Hello world!"
2,2022-05-12,"These messages are loaded from a CSV file."
3,2022-05-13,"Incredible!"`)

	builder := &strings.Builder{}
	table := New(builder)
	if err := table.LoadCSV(input, true); err != nil {
		panic(err)
	}
	table.Render()

	assertMultilineEqual(t, `
┌────┬────────────┬────────────────────────────────────────────┐
│ Id │    Date    │                  Message                   │
├────┼────────────┼────────────────────────────────────────────┤
│ 1  │ 2022-05-12 │ Hello world!                               │
├────┼────────────┼────────────────────────────────────────────┤
│ 2  │ 2022-05-12 │ These messages are loaded from a CSV file. │
├────┼────────────┼────────────────────────────────────────────┤
│ 3  │ 2022-05-13 │ Incredible!                                │
└────┴────────────┴────────────────────────────────────────────┘
`, "\n"+builder.String())
}

func Test_MultipleHeaderRows(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddHeaders("D", "E", "F")
	table.AddRow("1", "22", "333")
	table.AddRow("4444", "55555", "666666")
	table.Render()
	assertMultilineEqual(t, `
┌──────┬───────┬────────┐
│  A   │   B   │   C    │
├──────┼───────┼────────┤
│  D   │   E   │   F    │
├──────┼───────┼────────┤
│ 1    │ 22    │ 333    │
├──────┼───────┼────────┤
│ 4444 │ 55555 │ 666666 │
└──────┴───────┴────────┘
`, "\n"+builder.String())
}

func Test_MultipleFooterRows(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddHeaders("D", "E", "F")
	table.AddRow("1", "22", "333")
	table.AddRow("4444", "55555", "666666")
	table.SetFooters("G", "H", "I")
	table.AddFooters("J", "K", "L")
	table.Render()
	assertMultilineEqual(t, `
┌──────┬───────┬────────┐
│  A   │   B   │   C    │
├──────┼───────┼────────┤
│  D   │   E   │   F    │
├──────┼───────┼────────┤
│ 1    │ 22    │ 333    │
├──────┼───────┼────────┤
│ 4444 │ 55555 │ 666666 │
├──────┼───────┼────────┤
│  G   │   H   │   I    │
├──────┼───────┼────────┤
│  J   │   K   │   L    │
└──────┴───────┴────────┘
`, "\n"+builder.String())
}

func Test_HeaderColSpan(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B & C")
	table.SetHeaderColSpans(0, 1, 2)
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬───────┐
│ A │ B & C │
├───┼───┬───┤
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
└───┴───┴───┘
`, "\n"+builder.String())
}

func Test_HeaderColSpanLargerHeading(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "This is a long heading")
	table.SetHeaderColSpans(0, 1, 2)
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬────────────────────────┐
│ A │ This is a long heading │
├───┼───────────┬────────────┤
│ 1 │ 2         │ 3          │
├───┼───────────┼────────────┤
│ 4 │ 5         │ 6          │
└───┴───────────┴────────────┘
`, "\n"+builder.String())
}

func Test_HeaderColSpanSmallerHeading(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B")
	table.SetHeaderColSpans(0, 1, 2)
	table.AddRow("1", "2", "This is some long data")
	table.AddRow("4", "5", "6")
	table.Render()
	assertMultilineEqual(t, `
┌───┬────────────────────────────┐
│ A │             B              │
├───┼───┬────────────────────────┤
│ 1 │ 2 │ This is some long data │
├───┼───┼────────────────────────┤
│ 4 │ 5 │ 6                      │
└───┴───┴────────────────────────┘
`, "\n"+builder.String())
}

func Test_HeaderColSpanTrivyKubernetesStyle(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("Namespace", "Resource", "Vulnerabilities", "Misconfigurations")
	table.AddHeaders("Namespace", "Resource", "Critical", "High", "Critical", "High")
	table.SetHeaderColSpans(0, 1, 1, 2, 2)
	table.SetAutoMergeHeaders(true)
	table.AddRow("default", "Deployment/app", "2", "5", "0", "3")
	table.AddRow("default", "Ingress/test", "-", "-", "1", "0")
	table.AddRow("default", "Service/test", "0", "0", "3", "0")
	table.Render()
	assertMultilineEqual(t, `
┌───────────┬────────────────┬─────────────────┬───────────────────┐
│ Namespace │    Resource    │ Vulnerabilities │ Misconfigurations │
│           │                ├──────────┬──────┼───────────┬───────┤
│           │                │ Critical │ High │ Critical  │ High  │
├───────────┼────────────────┼──────────┼──────┼───────────┼───────┤
│ default   │ Deployment/app │ 2        │ 5    │ 0         │ 3     │
├───────────┼────────────────┼──────────┼──────┼───────────┼───────┤
│ default   │ Ingress/test   │ -        │ -    │ 1         │ 0     │
├───────────┼────────────────┼──────────┼──────┼───────────┼───────┤
│ default   │ Service/test   │ 0        │ 0    │ 3         │ 0     │
└───────────┴────────────────┴──────────┴──────┴───────────┴───────┘
`, "\n"+builder.String())
}

func Test_HeaderColSpanTrivyKubernetesStyleFull(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("Namespace", "Resource", "Vulnerabilities", "Misconfigurations")
	table.AddHeaders("Namespace", "Resource", "Critical", "High", "Medium", "Low", "Unknown", "Critical", "High", "Medium", "Low", "Unknown")
	table.SetHeaderColSpans(0, 1, 1, 5, 5)
	table.SetAutoMergeHeaders(true)
	table.AddRow("default", "Deployment/app", "2", "5", "7", "8", "0", "0", "3", "5", "19", "0")
	table.AddRow("default", "Ingress/test", "-", "-", "-", "-", "-", "1", "0", "2", "17", "0")
	table.AddRow("default", "Service/test", "0", "0", "0", "1", "0", "3", "0", "4", "9", "0")
	table.Render()
	assertMultilineEqual(t, `
┌───────────┬────────────────┬──────────────────────────────────────────┬──────────────────────────────────────────┐
│ Namespace │    Resource    │             Vulnerabilities              │            Misconfigurations             │
│           │                ├──────────┬──────┬────────┬─────┬─────────┼──────────┬──────┬────────┬─────┬─────────┤
│           │                │ Critical │ High │ Medium │ Low │ Unknown │ Critical │ High │ Medium │ Low │ Unknown │
├───────────┼────────────────┼──────────┼──────┼────────┼─────┼─────────┼──────────┼──────┼────────┼─────┼─────────┤
│ default   │ Deployment/app │ 2        │ 5    │ 7      │ 8   │ 0       │ 0        │ 3    │ 5      │ 19  │ 0       │
├───────────┼────────────────┼──────────┼──────┼────────┼─────┼─────────┼──────────┼──────┼────────┼─────┼─────────┤
│ default   │ Ingress/test   │ -        │ -    │ -      │ -   │ -       │ 1        │ 0    │ 2      │ 17  │ 0       │
├───────────┼────────────────┼──────────┼──────┼────────┼─────┼─────────┼──────────┼──────┼────────┼─────┼─────────┤
│ default   │ Service/test   │ 0        │ 0    │ 0      │ 1   │ 0       │ 3        │ 0    │ 4      │ 9   │ 0       │
└───────────┴────────────────┴──────────┴──────┴────────┴─────┴─────────┴──────────┴──────┴────────┴─────┴─────────┘
`, "\n"+builder.String())
}

func Test_RelativeColIndexesSimple(t *testing.T) {

	row := iRow{
		cols: []iCol{
			{
				span: 1,
			},
			{
				span: 1,
			},
			{
				span: 1,
			},
		},
	}

	table := New(nil)
	assert.Equal(t, 0, table.getRealIndex(row, 0))
	assert.Equal(t, 1, table.getRealIndex(row, 1))
	assert.Equal(t, 2, table.getRealIndex(row, 2))
	assert.Equal(t, 0, table.getRelativeIndex(row, 0))
	assert.Equal(t, 1, table.getRelativeIndex(row, 1))
	assert.Equal(t, 2, table.getRelativeIndex(row, 2))

}

func Test_RelativeColIndexesWithSpans(t *testing.T) {

	row := iRow{
		cols: []iCol{
			{
				span: 2,
			},
			{
				span: 3,
			},
			{
				span: 1,
			},
		},
	}

	table := New(nil)
	assert.Equal(t, 1, table.getRealIndex(row, 2))
	assert.Equal(t, 0, table.getRealIndex(row, 0))
	assert.Equal(t, 2, table.getRealIndex(row, 5))

	assert.Equal(t, 0, table.getRelativeIndex(row, 0))
	assert.Equal(t, 2, table.getRelativeIndex(row, 1))
	assert.Equal(t, 5, table.getRelativeIndex(row, 2))
}

func Test_HeaderColSpanVariation(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("Service", "Misconfigurations", "Last Scanned")
	table.AddHeaders("Service", "Critical", "High", "Medium", "Low", "Unknown", "Last Scanned")
	table.SetRowLines(false)
	table.SetHeaderAlignment(AlignLeft, AlignCenter, AlignCenter, AlignCenter, AlignCenter, AlignCenter, AlignLeft)
	table.SetAlignment(AlignLeft, AlignRight, AlignRight, AlignRight, AlignRight, AlignRight, AlignLeft)
	table.SetAutoMergeHeaders(true)
	table.SetHeaderColSpans(0, 1, 5, 1)
	table.AddRow("ec2", "1", "2", "5", "0", "3", "2 hours ago")
	table.AddRow("ecs", "0", "-", "-", "1", "0", "just now")
	table.AddRow("eks", "7", "0", "0", "3", "0", "127 hours ago")
	table.Render()
	assertMultilineEqual(t, `
┌─────────┬──────────────────────────────────────────────────┬───────────────┐
│ Service │                Misconfigurations                 │ Last Scanned  │
│         ├──────────┬──────────────┬────────┬─────┬─────────┤               │
│         │ Critical │     High     │ Medium │ Low │ Unknown │               │
├─────────┼──────────┼──────────────┼────────┼─────┼─────────┼───────────────┤
│ ec2     │        1 │            2 │      5 │   0 │       3 │ 2 hours ago   │
│ ecs     │        0 │            - │      - │   1 │       0 │ just now      │
│ eks     │        7 │            0 │      0 │   3 │       0 │ 127 hours ago │
└─────────┴──────────┴──────────────┴────────┴─────┴─────────┴───────────────┘
`, "\n"+builder.String())
}

func Test_HeaderVerticalAlignBottom(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("Service", "Misconfigurations", "Last Scanned")
	table.AddHeaders("Service", "Critical", "High", "Medium", "Low", "Unknown", "Last Scanned")
	table.SetRowLines(false)
	table.SetHeaderAlignment(AlignLeft, AlignCenter, AlignCenter, AlignCenter, AlignCenter, AlignCenter, AlignLeft)
	table.SetHeaderVerticalAlignment(AlignBottom)
	table.SetAlignment(AlignLeft, AlignRight, AlignRight, AlignRight, AlignRight, AlignRight, AlignLeft)
	table.SetAutoMergeHeaders(true)
	table.SetHeaderColSpans(0, 1, 5, 1)
	table.AddRow("ec2", "1", "2", "5", "0", "3", "2 hours ago")
	table.AddRow("ecs", "0", "-", "-", "1", "0", "just now")
	table.AddRow("eks", "7", "0", "0", "3", "0", "127 hours ago")
	table.Render()
	assertMultilineEqual(t, `
┌─────────┬──────────────────────────────────────────────────┬───────────────┐
│         │                Misconfigurations                 │               │
│         ├──────────┬──────────────┬────────┬─────┬─────────┤               │
│ Service │ Critical │     High     │ Medium │ Low │ Unknown │ Last Scanned  │
├─────────┼──────────┼──────────────┼────────┼─────┼─────────┼───────────────┤
│ ec2     │        1 │            2 │      5 │   0 │       3 │ 2 hours ago   │
│ ecs     │        0 │            - │      - │   1 │       0 │ just now      │
│ eks     │        7 │            0 │      0 │   3 │       0 │ 127 hours ago │
└─────────┴──────────┴──────────────┴────────┴─────┴─────────┴───────────────┘
`, "\n"+builder.String())
}

func Test_FillWidth(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.SetAvailableWidth(19)
	table.SetFillWidth(true)
	table.Render()
	assertMultilineEqual(t, `
┌─────┬─────┬─────┐
│  A  │  B  │  C  │
├─────┼─────┼─────┤
│ 1   │ 2   │ 3   │
├─────┼─────┼─────┤
│ 4   │ 5   │ 6   │
└─────┴─────┴─────┘
`, "\n"+builder.String())
}

func Test_HeaderColSpanTrivyKubernetesStyleFullWithFillWidth(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("Namespace", "Resource", "Vulnerabilities", "Misconfigurations")
	table.AddHeaders("Namespace", "Resource", "Critical", "High", "Medium", "Low", "Unknown", "Critical", "High", "Medium", "Low", "Unknown")
	table.SetHeaderColSpans(0, 1, 1, 5, 5)
	table.SetAutoMergeHeaders(true)
	table.SetAvailableWidth(100)
	table.SetFillWidth(true)
	table.AddRow("default", "Deployment/app", "2", "5", "7", "8", "0", "0", "3", "5", "19", "0")
	table.AddRow("default", "Ingress/test", "-", "-", "-", "-", "-", "1", "0", "2", "17", "0")
	table.AddRow("default", "Service/test", "0", "0", "0", "1", "0", "3", "0", "4", "9", "0")
	table.Render()
	assertMultilineEqual(t, `
┌──────────────┬──────────────────┬──────────────────────────────────────────┬───────────────────────────────────────────┐
│  Namespace   │     Resource     │             Vulnerabilities              │             Misconfigurations             │
│              │                  ├──────────┬──────┬────────┬─────┬─────────┼──────────┬──────┬────────┬──────┬─────────┤
│              │                  │ Critical │ High │ Medium │ Low │ Unknown │ Critical │ High │ Medium │ Low  │ Unknown │
├──────────────┼──────────────────┼──────────┼──────┼────────┼─────┼─────────┼──────────┼──────┼────────┼──────┼─────────┤
│ default      │ Deployment/app   │ 2        │ 5    │ 7      │ 8   │ 0       │ 0        │ 3    │ 5      │ 19   │ 0       │
├──────────────┼──────────────────┼──────────┼──────┼────────┼─────┼─────────┼──────────┼──────┼────────┼──────┼─────────┤
│ default      │ Ingress/test     │ -        │ -    │ -      │ -   │ -       │ 1        │ 0    │ 2      │ 17   │ 0       │
├──────────────┼──────────────────┼──────────┼──────┼────────┼─────┼─────────┼──────────┼──────┼────────┼──────┼─────────┤
│ default      │ Service/test     │ 0        │ 0    │ 0      │ 1   │ 0       │ 3        │ 0    │ 4      │ 9    │ 0       │
└──────────────┴──────────────────┴──────────┴──────┴────────┴─────┴─────────┴──────────┴──────┴────────┴──────┴─────────┘
`, "\n"+builder.String())
}

func Test_TableIsEmpty(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	assert.Equal(t, true, table.IsEmpty())
}

func Test_TableIsNotEmpty(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	assert.Equal(t, false, table.IsEmpty())
}

func Test_TableRowCount(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	assert.Equal(t, 0, table.RowCount())
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	assert.Equal(t, 2, table.RowCount())
}

func Test_ClearTable(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	assert.Equal(t, true, table.IsEmpty())
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	assert.Equal(t, false, table.IsEmpty())
	table.Clear()
	assert.Equal(t, true, table.IsEmpty())
}
