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

	var total int

	for _, line := range lines {
		if line == "" {
			continue
		}

		result := FindHighestCombination(line)
		total += result

		fmt.Println(line, "=", result)

	}

	fmt.Println("Sum =", total)
}

func readInput() []string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return slices.Collect(strings.SplitSeq(string(input), "\n"))
}

func FindHighestCombination(input string) int {
	var result string

	for i := range 12 {
		v := findHighest(input, 12-(i+1))
		result += v

		index := strings.Index(input, v)
		input = input[index+1:]
	}

	out, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}

	return out
}

func findHighest(input string, order int) string {
	var highest byte

	for i := range len(input) - order {
		if input[i] > highest {
			highest = input[i]
		}
	}

	if highest < '0' || highest > '9' {
		panic("unexpected value")
	}

	return string(highest)
}
