package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CSIStrip(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "hello world!",
			want:  "hello world!",
		},
		{
			input: "\x1b[37mhello \x1b[31mworld!",
			want:  "hello world!",
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, newANSI(test.input).Strip())
		})
	}
}

func Test_ANSILen(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "hello world!",
			want:  12,
		},
		{
			input: "\x1b[37mhello \x1b[31mworld!",
			want:  12,
		},
		{
			input: "🔥 unicode 🔥 characters 🔥",
			want:  27,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, newANSI(test.input).Len())
		})
	}
}
