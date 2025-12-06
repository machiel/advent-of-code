package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Merge(t *testing.T) {
	// 	3-5
	// 10-14
	// 16-20
	// 12-18
	input := []Range{
		{Start: 10, End: 14},
		{Start: 16, End: 20},
		{Start: 3, End: 5},
		{Start: 12, End: 18},
		{Start: 4, End: 4},
	}

	expected := []Range{
		{Start: 3, End: 5},
		{Start: 10, End: 20},
	}

	result := merge(input)
	assert.Equal(t, expected, result)
}

func TestRange_Length(t *testing.T) {
	tests := []struct {
		Range    Range
		Expected int
	}{
		{
			Range:    Range{Start: 0, End: 2},
			Expected: 3,
		},
		{
			Range:    Range{Start: 1, End: 2},
			Expected: 2,
		},
		{
			Range:    Range{Start: 2, End: 2},
			Expected: 1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			length := test.Range.Length()

			assert.Equal(t, test.Expected, length)
		})
	}
}
