package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseMap(t *testing.T) {
	input := []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}

	expected := []Map{
		{7},
		{6, 8},
		{5, 7, 9},
		{4, 6, 10},
		{3, 5, 9, 11},
		{2, 6, 12},
		{1, 3, 5, 7, 9, 13},
	}

	assert.Equal(t, expected, parseMap(slices.Values(input)))
}

func Test_lineToMap(t *testing.T) {
	tests := map[string]Map{
		"..^..^..S": {2, 5},
	}

	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			actual := lineToMap(input)

			assert.Equal(t, expected, actual)
		})
	}
}

var input = []Map{
	{7},
	{6, 8},
	{5, 7, 9},
	{4, 6, 10},
	{3, 5, 9, 11},
	{2, 6, 12},
	{1, 3, 5, 7, 9, 13},
}

func Test_buildGraph(t *testing.T) {
	actual := buildGraph(input)
	assert.Equal(t, adventGraph(), actual)
}

func Test_countPathways(t *testing.T) {
	actual := countPathways(adventGraph())
	assert.Equal(t, 40, actual)
}

func adventGraph() *Node {
	y7x1 := n(1, nil, nil)
	y7x3 := n(3, nil, nil)
	y7x5 := n(5, nil, nil)
	y7x7 := n(7, nil, nil)
	y7x13 := n(13, nil, nil)

	y6x2 := n(2, y7x1, y7x3)
	y6x6 := n(6, y7x5, y7x7)
	y6x12 := n(12, nil, y7x13)

	y5x3 := n(3, y6x2, nil)
	y5x5 := n(5, nil, y6x6)
	y5x9 := n(9, nil, nil)
	y5x11 := n(11, nil, y6x12)

	y4x4 := n(4, y5x3, y5x5)
	y4x6 := n(6, y5x5, y7x7)
	y4x10 := n(10, y5x9, y5x11)

	y3x5 := n(5, y4x4, y4x6)
	y3x7 := n(7, y4x6, nil)
	y3x9 := n(9, nil, y4x10)

	y2x6 := n(6, y3x5, y3x7)
	y2x8 := n(8, y3x7, y3x9)

	return n(7, y2x6, y2x8)
}

func n(index int, l, r *Node) *Node {
	return &Node{
		Index: Position(index),
		Left:  l,
		Right: r,
	}
}

func Test_countPathways2(t *testing.T) {
	tests := map[string]struct {
		Input    *Node
		Expected int
	}{
		"nil": {
			Input:    nil,
			Expected: 1,
		},
		"leaf_node": {
			Input:    &Node{},
			Expected: 2,
		},
		"node_with_one_leaf_left": {
			Input: &Node{
				Left: &Node{},
			},
			Expected: 3,
		},
		"node_with_one_leaf_right": {
			Input: &Node{
				Right: &Node{},
			},
			Expected: 3,
		},
		"node_with_two_leaves": {
			Input: &Node{
				Left:  &Node{},
				Right: &Node{},
			},
			Expected: 4,
		},
		"node_with_node_with_two_leaves": {
			Input: &Node{
				Left: &Node{
					Left:  &Node{},
					Right: &Node{},
				},
			},
			Expected: 5,
		},
		"from_11_downwards": {
			Input: &Node{
				Index: 11,
				Right: &Node{
					Index: 12,
					Right: &Node{
						Index: 13,
					},
				},
			},
			Expected: 4,
		},
		"from_10_downwards": {
			Input: &Node{
				Index: 10,
				Left: &Node{
					Index: 9,
				},
				Right: &Node{
					Index: 11,
					Right: &Node{
						Index: 12,
						Right: &Node{
							Index: 13,
						},
					},
				},
			},
			Expected: 6,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := countPathways(test.Input)
			assert.Equal(t, test.Expected, actual)
		})
	}
}
