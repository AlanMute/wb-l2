package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_SortByColumn(t *testing.T) {
	var table = []struct {
		inputData   [][]string
		k           int
		n           bool
		expectedOut [][]string
	}{
		{
			inputData: [][]string{
				{"a", "b", "c"},
				{"b", "a", "b"},
				{"c", "a", "b"},
			},
			k: 1,
			n: false,
			expectedOut: [][]string{
				{"a", "b", "c"},
				{"b", "a", "b"},
				{"c", "a", "b"},
			},
		},
		{
			inputData: [][]string{
				{"a", "b", "c"},
				{"b", "a", "b"},
				{"c", "z", "b"},
			},
			k: 2,
			n: false,
			expectedOut: [][]string{
				{"b", "a", "b"},
				{"a", "b", "c"},
				{"c", "z", "b"},
			},
		},
		{
			inputData: [][]string{
				{"11", "b", "c"},
				{"12", "a", "b"},
				{"2", "z", "b"},
				{"a", "b"},
			},
			k: 1,
			n: true,
			expectedOut: [][]string{
				{"2", "z", "b"},
				{"11", "b", "c"},
				{"12", "a", "b"},
				{"a", "b"},
			},
		},
		{
			inputData: [][]string{
				{"a", "b", "c"},
				{"b", "a", "b"},
				{"c", "z", "b"},
			},
			k: 4,
			n: false,
			expectedOut: [][]string{
				{"a", "b", "c"},
				{"b", "a", "b"},
				{"c", "z", "b"},
			},
		},
	}

	for _, test := range table {
		output := SortByColumn(test.inputData, test.k, test.n)

		assert.Equal(t, test.expectedOut, output)
	}
}

func Test_ReverseSlice(t *testing.T) {
	var table = []struct {
		inputData   [][]string
		expectedOut [][]string
	}{
		{
			inputData: [][]string{
				{"a", "b", "c"},
				{"b", "a", "b"},
				{"c", "a", "b"},
			},
			expectedOut: [][]string{
				{"c", "a", "b"},
				{"b", "a", "b"},
				{"a", "b", "c"},
			},
		},
		{
			inputData: [][]string{
				{"a", "b"},
			},
			expectedOut: [][]string{
				{"a", "b"},
			},
		},
	}

	for _, test := range table {
		output := ReverseSlice(test.inputData)

		assert.Equal(t, test.expectedOut, output)
	}
}

func Test_UniqueStrings(t *testing.T) {
	var table = []struct {
		inputData   [][]string
		expectedOut [][]string
	}{
		{
			inputData: [][]string{
				{"a", "b", "c"},
				{"a", "b", "c"},
				{"c", "a", "b"},
			},
			expectedOut: [][]string{
				{"a", "b", "c"},
				{"c", "a", "b"},
			},
		},
		{
			inputData: [][]string{
				{"a", "b"},
			},
			expectedOut: [][]string{
				{"a", "b"},
			},
		},
	}

	for _, test := range table {
		output := UniqueStrings(test.inputData)

		assert.Equal(t, test.expectedOut, output)
	}
}
