package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getAccessibleRolls(t *testing.T) {
	grid := [][]string{
		{".", ".", "@", "@", ".", "@", "@", "@", "@", "."},
		{"@", "@", "@", ".", "@", ".", "@", ".", "@", "@"},
		{"@", "@", "@", "@", "@", ".", "@", ".", "@", "@"},
		{"@", ".", "@", "@", "@", "@", ".", ".", "@", "."},
		{"@", "@", ".", "@", "@", "@", "@", ".", "@", "@"},
		{".", "@", "@", "@", "@", "@", "@", "@", ".", "@"},
		{".", "@", ".", "@", ".", "@", ".", "@", "@", "@"},
		{"@", ".", "@", "@", "@", ".", "@", "@", "@", "@"},
		{".", "@", "@", "@", "@", "@", "@", "@", "@", "."},
		{"@", ".", "@", ".", "@", "@", "@", ".", "@", "."},
	}

	assert.Equal(t, 13, getAccessibleRolls(grid))
}

func Test_getAdjectCount(t *testing.T) {
	grid := [][]string{
		{".", ".", "@", "@", ".", "@", "@", "@", "@", "."},
		{"@", "@", "@", ".", "@", ".", "@", ".", "@", "@"},
		{"@", "@", "@", "@", "@", ".", "@", ".", "@", "@"},
		{"@", ".", "@", "@", "@", "@", ".", ".", "@", "."},
		{"@", "@", ".", "@", "@", "@", "@", ".", "@", "@"},
		{".", "@", "@", "@", "@", "@", "@", "@", ".", "@"},
		{".", "@", ".", "@", ".", "@", ".", "@", "@", "@"},
		{"@", ".", "@", "@", "@", ".", "@", "@", "@", "@"},
		{".", "@", "@", "@", "@", "@", "@", "@", "@", "."},
		{"@", ".", "@", ".", "@", "@", "@", ".", "@", "."},
	}

	tests := map[string]struct {
		Row           int
		Column        int
		ExpectedCount int
	}{
		"0_0": {
			Row:           0,
			Column:        0,
			ExpectedCount: 0,
		},
		"0_2": {
			Row:           0,
			Column:        2,
			ExpectedCount: 3,
		},
		"3_0": {
			Row:           3,
			Column:        0,
			ExpectedCount: 4,
		},
		"2_3": {
			Row:           2,
			Column:        3,
			ExpectedCount: 7,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			count := getAdjacentCount(grid, test.Row, test.Column)
			assert.Equal(t, test.ExpectedCount, count)
		})
	}
}

func Test_hasRoll(t *testing.T) {
	grid := [][]string{
		{".", "@", "."},
		{"@", "@", "@"},
		{"@", ".", "@"},
	}

	tests := map[string]struct {
		Row      int
		Column   int
		Expected bool
	}{
		"has_roll": {
			Row:      0,
			Column:   1,
			Expected: true,
		},
		"does_not_have_roll": {
			Row:    0,
			Column: 0,
		},
		"has_roll_last_row_and_column": {
			Row:      2,
			Column:   2,
			Expected: true,
		},
		"row_out_of_bounds": {
			Row: -1,
		},
		"row_out_of_bounds_2": {
			Row: 3,
		},
		"column_out_of_bounds": {
			Column: -1,
		},
		"column_out_of_bounds_2": {
			Column: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := hasRoll(grid, test.Row, test.Column)
			assert.Equal(t, test.Expected, actual)
		})
	}
}
