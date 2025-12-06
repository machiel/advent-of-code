package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()

	var lefts []int
	rights := make(map[int]int, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}

		left, right := parseLine(line)

		lefts = append(lefts, left)
		rights[right]++
	}

	var total int

	for _, v := range lefts {
		score := rights[v] * v
		total += score
	}

	fmt.Println(total)
}

func partOne() {
	lines := readInput()

	var lefts []int
	var rights []int

	for _, line := range lines {
		if line == "" {
			continue
		}

		left, right := parseLine(line)

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	var diff int

	for k, v := range lefts {
		diff += int(math.Abs(float64(rights[k] - v)))
	}

	fmt.Println(diff)
}

func readInput() []string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return slices.Collect(strings.SplitSeq(string(input), "\n"))
}

func parseLine(in string) (int, int) {
	parts := strings.Fields(in)

	if len(parts) != 2 {
		panic("invalid input " + in)
	}

	return parse(parts[0]), parse(parts[1])
}

func parse(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return i
}
