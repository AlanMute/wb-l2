package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_manGrep(t *testing.T) {
	var table = []struct {
		input       string
		words       string
		i           bool //игнорировать регистр
		v           bool //вместо совпадения, исключать
		F           bool //точное совпадение со строкой, не паттерн
		expectedOut bool
	}{
		{
			input:       "aaa bbb ccc ddd",
			words:       "aaa",
			i:           false,
			v:           false,
			F:           false,
			expectedOut: true,
		},
		{
			input:       "",
			words:       "",
			i:           false,
			v:           false,
			F:           false,
			expectedOut: true,
		},
		{
			input:       "aaa bbb ccc ddd",
			words:       "aAa",
			i:           true,
			v:           false,
			F:           false,
			expectedOut: true,
		},
		{
			input:       "aaa bbb ccc ddd",
			words:       "aAa",
			i:           true,
			v:           true,
			F:           false,
			expectedOut: false,
		},
		{
			input:       "aaa bbb ccc ddd",
			words:       "aAa",
			i:           true,
			v:           false,
			F:           true,
			expectedOut: false,
		},
		{
			input:       "aaa bbb ccc ddd",
			words:       "aaa Bbb ccC ddd",
			i:           true,
			v:           false,
			F:           true,
			expectedOut: true,
		},
	}

	for _, test := range table {
		output := manGrep(test.input, test.words, test.i, test.v, test.F)

		assert.Equal(t, test.expectedOut, output)
	}
}
