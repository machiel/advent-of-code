package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHighestCombination(t *testing.T) {
	tests := map[string]int{
		"987654321111111": 987654321111,
		"811111111111119": 811111111119,
		"234234234234278": 434234234278,
		"818181911112111": 888911112111,
	}

	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			actual := FindHighestCombination(input)

			assert.Equal(t, expected, actual)
		})
	}
}

func Test_findHighest(t *testing.T) {
	tests := map[string]struct {
		Expected string
		Order    int
	}{
		"987654321111111": {
			Expected: "9",
			Order:    1,
		},
		"811111111111119": {
			Expected: "8",
			Order:    1,
		},
		"234234234234278": {
			Expected: "7",
			Order:    1,
		},
		"818181911112111": {
			Expected: "9",
			Order:    1,
		},
		"234234234234268": {
			Expected: "8",
			Order:    0,
		},
	}

	for input, test := range tests {
		t.Run(input, func(t *testing.T) {
			actual := findHighest(input, test.Order)

			assert.Equal(t, test.Expected, actual)
		})
	}
}
