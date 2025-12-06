package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()

	ranges := make([]Range, 0)

	for line := range lines {
		if line == "" {
			continue
		}

		r, err := parseRange(line)
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, r)
	}

	var total int

	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			if isRepetitive(i) {
				total += i
			}
		}
	}

	fmt.Println("Result:", total)
}

func readInput() iter.Seq[string] {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	in := strings.ReplaceAll(string(input), "\n", "")
	in = strings.ReplaceAll(in, " ", "")

	return strings.SplitSeq(in, ",")
}

func parseRange(in string) (Range, error) {
	parts := strings.Split(in, "-")

	if len(parts) != 2 {
		return Range{}, fmt.Errorf("invalid input: %s", in)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return Range{}, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return Range{}, err
	}

	return Range{
		Start: start,
		End:   end,
	}, nil
}

type Range struct {
	Start int
	End   int
}

func isRepetitive(i int) bool {
	return FindRepetive(strconv.Itoa(i))
}

func FindRepetive(input string) bool {
	l := len(input)

	for i := 1; i < len(input); i++ {
		if l%i != 0 {
			continue
		}

		repeats := l / i
		comparePart := input[:i]
		compareString := strings.Repeat(comparePart, repeats)

		if input == compareString {
			return true
		}
	}

	return false
}
