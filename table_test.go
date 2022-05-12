package table

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BasicTable(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRow("1", "2", "3")
	table.AddRow("4", "5", "6")
	table.Render()
	assert.Equal(t, `
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
	assert.Equal(t, ``, builder.String())
}

func Test_AddRows(t *testing.T) {
	builder := &strings.Builder{}
	table := New(builder)
	table.SetHeaders("A", "B", "C")
	table.AddRows([]string{"1", "2", "3"}, []string{"4", "5", "6"})
	table.Render()
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
┌───┬───┬───┐
│ 1 │ 2 │ 3 │
├───┼───┼───┤
│ 4 │ 5 │ 6 │
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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
	assert.Equal(t, `
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

	assert.Equal(t, `
┌─────────────────────────────┬───┬───┐
│              A              │ B │ C │
├─────────────────────────────┼───┼───┤
│ 🔥 unicode 🔥 characters 🔥 │ 2 │ 3 │
├─────────────────────────────┼───┼───┤
│ 4                           │ 5 │ 6 │
└─────────────────────────────┴───┴───┘
`, "\n"+builder.String())
}
