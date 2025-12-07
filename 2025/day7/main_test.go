package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evaluate(t *testing.T) {
	m := []Map{
		{Start: toPtr(Position(3))},
		{Splitters: []Position{2, 4}},
		{Splitters: []Position{3}},
		{Splitters: []Position{2}},
	}

	totalTimelines := evaluate(World{}, m)
	assert.Equal(t, totalTimelines, 3)
}

func Test_lineToMap(t *testing.T) {
	tests := map[string]Map{
		"..S..": {
			Start: toPtr(Position(2)),
		},
		"..^..^..S": {
			Start:     toPtr(Position(8)),
			Splitters: []Position{2, 5},
		},
	}

	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			actual := lineToMap(input)

			assert.Equal(t, expected, actual)
		})
	}
}
