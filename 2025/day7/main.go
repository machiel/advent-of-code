package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strings"
)

func main() {
	var maps []Map

	for line := range readInput() {
		m := lineToMap(line)

		maps = append(maps, m)
	}

	fmt.Println(evaluate(World{}, maps))
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
	Beam Position
}

var cache = map[World]map[string]int{}

func evaluate(w World, m []Map) int {
	if len(m) == 0 {
		return 1
	}

	entry := m[0]

	if v, ok := cache[w][entry.Line]; ok {
		return v
	}

	if entry.Start != nil {
		return evaluate(World{
			Beam: *entry.Start,
		}, m[1:])
	}

	var totalSplits int

	for _, s := range entry.Splitters {
		if s != w.Beam {
			continue
		}

		totalSplits += evaluate(World{Beam: s - 1}, m[1:])
		totalSplits += evaluate(World{Beam: s + 1}, m[1:])

		if _, ok := cache[w]; !ok {
			cache[w] = make(map[string]int)
		}

		cache[w][entry.Line] = totalSplits

		return totalSplits
	}

	return evaluate(w, m[1:])
}

type MapKey string

type Map struct {
	Line      string
	Start     *Position
	Splitters []Position
}

func lineToMap(line string) Map {
	m := Map{
		Line: line,
	}

	for i, v := range line {
		switch v {
		case 'S':
			m.Start = toPtr(Position(i))
		case '^':
			m.Splitters = append(m.Splitters, Position(i))
		case '.':
		default:
			panic("Unknown character: " + string(v))
		}
	}

	return m
}

func toPtr[T any](t T) *T {
	return &t
}
