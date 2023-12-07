package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_cut(t *testing.T) {
	var table = []struct {
		input     []string
		fields    int
		delimiter string
		separated bool
		expectOut []string
	}{
		{
			input: []string{
				"a b c",
				"d e f",
				"g h i",
			},
			fields:    1,
			delimiter: " ",
			separated: false,
			expectOut: []string{
				"a", "d", "g",
			},
		},
		{
			input: []string{
				"a b c",
				"d e f",
				"g h i",
			},
			fields:    2,
			delimiter: " ",
			separated: false,
			expectOut: []string{
				"b", "e", "h",
			},
		},
		{
			input: []string{
				"a b c",
				"d e f",
				"g h i",
			},
			fields:    4,
			delimiter: " ",
			separated: false,
		},
		{
			input: []string{
				"a b c",
				"d e f",
				"ghi",
			},
			fields:    1,
			delimiter: " ",
			separated: true,
			expectOut: []string{
				"a", "d",
			},
		},
	}

	for _, test := range table {
		output := cut(test.input, test.fields, test.delimiter, test.separated)

		assert.Equal(t, output, test.expectOut)
	}
}
