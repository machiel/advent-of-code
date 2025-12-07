package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strings"
)

func main() {
	n := buildGraph(parseMap(readInput()))
	fmt.Println("Finished building graph")

	fmt.Println(countPathways(n))
}

func parseMap(lines iter.Seq[string]) []Map {
	var maps []Map

	for line := range lines {
		m := lineToMap(line)

		if !m.Empty() {
			maps = append(maps, m)
		}
	}

	return maps
}

func readInput() iter.Seq[string] {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return strings.SplitSeq(string(input), "\n")
}

type Position int

type Map []Position

func (m Map) Empty() bool {
	return len(m) == 0
}

func lineToMap(line string) Map {
	var m Map

	for i, v := range line {
		switch v {
		case '^':
			m = append(m, Position(i))
		case 'S', '.':
		default:
			panic("Unknown character: " + string(v))
		}
	}

	return m
}

func buildGraph(in []Map) *Node {
	root := &Node{
		Index: in[0][0],
	}

	pendingNodes := map[Position][]*Node{
		in[0][0]: {root},
	}

	for _, m := range in[1:] {
		for _, pos := range m {

			n := &Node{
				Index: pos,
			}

			var foundParent bool

			if parents, ok := pendingNodes[pos-1]; ok {
				for _, parent := range parents {
					if parent.Right == nil {
						parent.Right = n

						foundParent = true
					}
				}
			}

			if parents, ok := pendingNodes[pos+1]; ok {
				for _, parent := range parents {
					if parent.Left == nil {
						parent.Left = n

						foundParent = true
					}
				}
			}

			if foundParent {
				pendingNodes[pos] = append(pendingNodes[pos], n)
			}
		}
	}

	return root
}

type Node struct {
	Index Position
	Left  *Node
	Right *Node

	Eval *int
}

func countPathways(n *Node) int {
	if n == nil {
		return 1
	}

	if n.Eval != nil {
		return *n.Eval
	}

	eval := countPathways(n.Left) + countPathways(n.Right)
	n.Eval = &eval
	return eval
}
