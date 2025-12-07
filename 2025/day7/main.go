package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strings"
)

func main() {
	fmt.Println(partOne(readInput()))
}

func readInput() iter.Seq[string] {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return strings.SplitSeq(string(input), "\n")
}

type Position int

type World struct {
	Beams map[Position]struct{}
}

func partOne(lines iter.Seq[string]) int {
	var world World
	var total int

	for line := range lines {
		m := lineToMap(line)

		var splits int
		world, splits = evaluate(world, m)

		total += splits
	}

	return total
}

func evaluate(w World, m Map) (World, int) {
	var splits int

	out := World{
		Beams: make(map[Position]struct{}, 0),
	}

	for _, s := range m.Starts {
		out.Beams[s] = struct{}{}
	}

	for _, s := range m.Splitters {
		_, ok := w.Beams[s]
		if !ok {
			continue
		}

		delete(w.Beams, s)

		out.Beams[s-1] = struct{}{}
		out.Beams[s+1] = struct{}{}

		splits++
	}

	for k := range w.Beams {
		out.Beams[k] = struct{}{}
	}

	return out, splits
}

type Map struct {
	Starts    []Position
	Splitters []Position
}

func lineToMap(line string) Map {
	var m Map

	for i, v := range line {
		switch v {
		case 'S':
			m.Starts = append(m.Starts, Position(i))
		case '^':
			m.Splitters = append(m.Splitters, Position(i))
		case '.':
		default:
			panic("Unknown character: " + string(v))
		}
	}

	return m
}
