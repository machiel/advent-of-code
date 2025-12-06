package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()

	ranges := make([]Range, 0)

	for _, line := range lines {
		if line == "" {
			break
		}

		r := parseRange(line)
		ranges = append(ranges, r)

		fmt.Println("Adding", r.Start, "to", r.End, "=", r.Length())
	}

	merged := merge(ranges)

	var total int
	for _, r := range merged {
		fmt.Println("Evaluating", r.Start, "to", r.End, "=", r.Length())
		total += r.Length()
	}
	fmt.Println(total)
}

func merge(in []Range) []Range {
	if len(in) == 0 {
		return in
	}

	slices.SortFunc(in, func(a, b Range) int {
		return a.Start - b.Start
	})

	out := make([]Range, 0)

	index := 1
	r := in[0]

	for index <= len(in)-1 {
		if !r.Contains(in[index].Start) {
			out = append(out, r)

			r = in[index]
			index++
			continue
		}

		if in[index].End > r.End {
			r.End = in[index].End
		}

		index++
		continue
	}

	out = append(out, r)

	return out
}

func firstPart() {
	lines := readInput()

	ranges := make([]Range, 0)

	var newLineCount int
	var matches int

	for _, line := range lines {
		if line == "" {
			newLineCount++
			continue
		}

		switch newLineCount {
		case 0:
			ranges = append(ranges, parseRange(line))
		case 1:
			id, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			var found bool

			for _, r := range ranges {
				if r.Contains(id) {
					fmt.Println(id, "matches", r)
					found = true
					break
				}
			}

			if !found {
				fmt.Println("No match for", id)
			} else {
				matches++
			}
		default:
		}
	}

	fmt.Println("Total matches", matches)
}

func readInput() []string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return slices.Collect(strings.SplitSeq(string(input), "\n"))
}

func parseRange(in string) Range {
	parts := strings.Split(in, "-")
	if len(parts) != 2 {
		fmt.Println(in)
		panic("Invalid input")
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return Range{
		Start: start,
		End:   end,
	}
}

type Range struct {
	Start int
	End   int
}

func (r Range) Overlaps(other Range) bool {
	return r.Contains(other.Start) || r.Contains(other.End)
}

func (r Range) Contains(i int) bool {
	return i >= r.Start && i <= r.End
}

func (r Range) Length() int {
	return (r.End - r.Start) + 1
}
