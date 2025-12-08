package main

import (
	"fmt"
	"io"
	"iter"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	positions := make([]*Position, 0)
	id := 1

	for line := range readInput() {
		if line == "" {
			continue
		}

		pos := parsePosition(id, line)

		positions = append(positions, &pos)
		id++
	}

	distances := make([]Distance, 0)

	for i, p := range positions {
		for _, q := range positions[i:] {
			if p == q {
				continue
			}

			distances = append(distances, Distance{
				Distance: p.Distance(q),
				Left:     p,
				Right:    q,
			})
		}
	}

	slices.SortFunc(distances, func(a, b Distance) int {
		return int(a.Distance - b.Distance)
	})

	t := NewTracker(positions)

	for _, d := range distances {
		t.Connect(d.Left, d.Right)

		if len(t.groups) == 1 {
			fmt.Println("WE DONE")
			fmt.Println(d.Left.X * d.Right.X)
			os.Exit(1)
		}
	}

	counts := map[int]int{}

	for k, v := range t.groups {
		counts[k] = len(v)
	}

	fmt.Println(counts)

	vals := slices.Collect(maps.Values(counts))
	slices.Sort(vals)
	fmt.Println(vals)
}

func readInput() iter.Seq[string] {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return strings.SplitSeq(string(input), "\n")
}

type Distance struct {
	Distance float64
	Left     *Position
	Right    *Position
}

type Position struct {
	ID int

	X int
	Y int
	Z int

	Group int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}

func (p Position) Distance(q *Position) float64 {
	return math.Sqrt(
		math.Pow(float64(p.X)-float64(q.X), 2) +
			math.Pow(float64(p.Y)-float64(q.Y), 2) +
			math.Pow(float64(p.Z)-float64(q.Z), 2))
}

func parsePosition(id int, in string) Position {
	parts := strings.Split(in, ",")
	if len(parts) != 3 {
		panic("Unexpected amount of parts")
	}

	return Position{
		ID: id,
		X:  parse(parts[0]),
		Y:  parse(parts[1]),
		Z:  parse(parts[2]),
	}
}

func parse(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return i
}

type Tracker struct {
	groups map[int][]*Position
}

func NewTracker(positions []*Position) *Tracker {
	groups := make(map[int][]*Position)

	for i, pos := range positions {
		pos.Group = i + 1
		groups[i+1] = []*Position{pos}
	}

	return &Tracker{
		groups: groups,
	}
}

func (t *Tracker) Connect(a, b *Position) {
	if a == b {
		return
	}

	if a.Group == b.Group {
		return
	}

	if a.Group == 0 {
		a.Group = b.Group
		t.groups[b.Group] = append(t.groups[b.Group], a)
		return
	}

	if b.Group == 0 {
		b.Group = a.Group
		t.groups[a.Group] = append(t.groups[a.Group], b)
		return
	}

	bGroup := b.Group

	for _, p := range t.groups[bGroup] {
		p.Group = a.Group
		t.groups[a.Group] = append(t.groups[a.Group], p)
	}

	delete(t.groups, bGroup)
}
