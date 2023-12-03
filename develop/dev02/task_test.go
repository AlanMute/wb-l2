package main

import (
	"fmt"
	"testing"
)

func Test_strUnpack(t *testing.T) {
	var table = []struct {
		input       string
		expectedOut string
		expectedErr error
	}{
		{
			input:       "a4bc2d5e",
			expectedOut: "aaaabccddddde",
			expectedErr: nil,
		},
		{
			input:       "abcd",
			expectedOut: "abcd",
			expectedErr: nil,
		},
		{
			input:       "45",
			expectedOut: "",
			expectedErr: fmt.Errorf("Некорректная строка!"),
		},
		{
			input:       "",
			expectedOut: "",
			expectedErr: nil,
		},
		{
			input:       "qwe\\4\\5",
			expectedOut: "qwe45",
			expectedErr: nil,
		},
		{
			input:       "qwe\\45",
			expectedOut: "qwe44444",
			expectedErr: nil,
		},
		{
			input:       "qwe\\\\5",
			expectedOut: "qwe\\\\\\\\\\",
			expectedErr: nil,
		},
	}

	for _, test := range table {
		output, err := strUnpack(test.input)

		if output != test.expectedOut || (err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error()) || (err != nil && test.expectedErr == nil) || (err == nil && test.expectedErr != nil) {
			t.Error("Wrong answer!")

			t.Errorf("\nInput %q\nOutput\t %q %s\nExpected %q %s", test.input, output, err.Error(), test.expectedOut, test.expectedErr.Error())
		}
	}
}
