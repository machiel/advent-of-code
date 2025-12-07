package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evaluate(t *testing.T) {
	tests := map[string]struct {
		InputWorld     World
		Map            Map
		ExpectedWorld  World
		ExpectedSplits int
	}{
		"start": {
			InputWorld: World{},
			Map: Map{
				Starts: []Position{2},
			},
			ExpectedWorld: World{
				Beams: map[Position]struct{}{
					2: {},
				},
			},
		},
		"beam_hits_split": {
			InputWorld: World{
				Beams: map[Position]struct{}{
					2: {},
				},
			},
			Map: Map{
				Splitters: []Position{2},
			},
			ExpectedWorld: World{
				Beams: map[Position]struct{}{
					1: {},
					3: {},
				},
			},
			ExpectedSplits: 1,
		},
		"beam_hits_split_and_one_doesnt": {
			InputWorld: World{
				Beams: map[Position]struct{}{
					2: {},
					5: {},
				},
			},
			Map: Map{
				Splitters: []Position{2},
			},
			ExpectedWorld: World{
				Beams: map[Position]struct{}{
					1: {},
					3: {},
					5: {},
				},
			},
			ExpectedSplits: 1,
		},
		"multi_split": {
			InputWorld: World{
				Beams: map[Position]struct{}{
					2: {},
					5: {},
				},
			},
			Map: Map{
				Splitters: []Position{2, 5},
			},
			ExpectedWorld: World{
				Beams: map[Position]struct{}{
					1: {},
					3: {},
					4: {},
					6: {},
				},
			},
			ExpectedSplits: 2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, actualSplits := evaluate(test.InputWorld, test.Map)

			assert.Equal(t, test.ExpectedWorld, actual)
			assert.Equal(t, test.ExpectedSplits, actualSplits)
		})
	}
}

func Test_lineToMap(t *testing.T) {
	tests := map[string]Map{
		"..S..": {
			Starts: []Position{2},
		},
		"..S.S": {
			Starts: []Position{2, 4},
		},
		"..^..^..S": {
			Starts:    []Position{8},
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
