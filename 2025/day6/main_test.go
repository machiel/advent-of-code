package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation_Solve(t *testing.T) {
	tests := map[string]struct {
		Input    Operation
		Expected int
	}{
		"add": {
			Input: Operation{
				Operation: Add,
				Values:    []int{1, 2, 3, 4},
			},
			Expected: 10,
		},
		"multiply": {
			Input: Operation{
				Operation: Multiply,
				Values:    []int{1, 2, 3, 4},
			},
			Expected: 24,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := test.Input.Solve()
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func Test_parseInput(t *testing.T) {
	input := []string{
		"100 20 4000 80 90",
		"  2  2  200 2  2 ",
		"3   3  30    3 33",
	}

	expected := [][]string{
		{"100", "20", "4000", "80", "90"},
		{"  2", " 2", " 200", "2 ", "2 "},
		{"3  ", "3 ", "30  ", " 3", "33"},
	}

	assert.Equal(t, expected, parseInput(slices.Values(input)))
}

func Test_resolveValues(t *testing.T) {
	tests := map[string]struct {
		Input  []string
		Output []int
	}{
		"one": {
			Input:  []string{"123", " 45", "  6"},
			Output: []int{356, 24, 1},
		},
		"two": {
			Input:  []string{"328", "64 ", "98 "},
			Output: []int{8, 248, 369},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Output, resolveValues(test.Input))
		})
	}
}

func Test_transform(t *testing.T) {
	in := [][]string{
		{"123", "328", " 51", "64 "},
		{" 45", "64 ", "387", "23 "},
		{"  6", "98 ", "215", "314"},
		{"*  ", "+  ", "*  ", "+  "},
	}

	expected := [][]string{
		{"123", " 45", "  6", "*  "},
		{"328", "64 ", "98 ", "+  "},
		{" 51", "387", "215", "*  "},
		{"64 ", "23 ", "314", "+  "},
	}

	assert.Equal(t, expected, transform(in))
}

func Test_parseOperations(t *testing.T) {
	in := [][]string{
		{"123", " 45", "  6", "*  "},
		{"328", "64 ", "98 ", "+  "},
		{" 51", "387", "215", "*  "},
		{"64 ", "23 ", "314", "+  "},
	}

	expected := []Operation{
		{
			Operation: Multiply,
			Values:    []int{356, 24, 1},
		},
		{
			Operation: Add,
			Values:    []int{8, 248, 369},
		},
		{
			Operation: Multiply,
			Values:    []int{175, 581, 32},
		},
		{
			Operation: Add,
			Values:    []int{4, 431, 623},
		},
	}

	actual := parseOperations(in)
	assert.Equal(t, expected, actual)
}
