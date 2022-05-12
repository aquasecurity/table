package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TextWrapping(t *testing.T) {
	tests := []struct {
		name  string
		input string
		wrap  int
		want  []string
	}{
		{
			name:  "no wrap required",
			input: "hello world!",
			wrap:  100,
			want:  []string{"hello world!"},
		},
		{
			name:  "basic wrap",
			input: "hello world!",
			wrap:  10,
			want:  []string{"hello", "world!"},
		},
		{
			name:  "exact length word",
			input: "hello incredible world!",
			wrap:  10,
			want:  []string{"hello", "incredible", "world!"},
		},
		{
			name:  "exact total fit",
			input: "it fit gud hooray",
			wrap:  10,
			want:  []string{"it fit gud", "hooray"},
		},
		{
			name:  "break word",
			input: "hello world, antidisestablishmentarianism!",
			wrap:  16,
			want:  []string{"hello world,", "antidisestablis-", "hmentarianism!"},
		},
		{
			name:  "new lines",
			input: "hello world\nthis is a\nlong sentence.",
			wrap:  10,
			want:  []string{"hello", "world", "this is a", "long", "sentence."},
		},
		{
			name:  "empty string",
			input: "",
			wrap:  10,
			want:  []string{""},
		},
		{
			name:  "multiple whitespace",
			input: "hello          world!",
			wrap:  10,
			want:  []string{"hello", "world!"},
		},
		{
			name:  "ansi codes",
			input: "\x1b[37mhello this should be\x1b[38mover 4 lines!",
			wrap:  10,
			want:  []string{"\x1b[37mhello this", "should", "be\x1b[38mover 4", "lines!"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := wrapText(test.input, test.wrap)
			assert.Equal(t, test.want, blobsToStrings(output))
		})
	}
}

func blobsToStrings(blobs []ansiBlob) []string {
	var output []string
	for _, blob := range blobs {
		output = append(output, blob.String())
	}
	return output
}
