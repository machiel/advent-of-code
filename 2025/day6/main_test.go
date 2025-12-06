package main

import (
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

func Test_transform(t *testing.T) {
	in := [][]string{
		{"123", "328", "51", "64"},
		{"45", "64", "387", "23"},
		{"6", "98", "215", "314"},
		{"*", "+", "*", "+"},
	}

	expected := []Operation{
		{
			Operation: Multiply,
			Values:    []int{123, 45, 6},
		},
		{
			Operation: Add,
			Values:    []int{328, 64, 98},
		},
		{
			Operation: Multiply,
			Values:    []int{51, 387, 215},
		},
		{
			Operation: Add,
			Values:    []int{64, 23, 314},
		},
	}

	actual := transform(in)
	assert.Equal(t, expected, actual)
}
