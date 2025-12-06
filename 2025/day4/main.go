package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	grid := readInput()

	var total int

	for {
		var count int
		grid, count = getAccessibleRolls(grid)

		fmt.Println("Removed", count, "rolls")

		total = total + count

		if count == 0 {
			break
		}
	}

	fmt.Println(total)
}

func readInput() [][]string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	result := make([][]string, 0)

	for line := range strings.SplitSeq(string(input), "\n") {
		out := make([]string, 0)

		for _, c := range line {
			out = append(out, string(c))
		}

		result = append(result, out)
	}

	return result
}

func getAccessibleRolls(grid [][]string) ([][]string, int) {
	out := make([][]string, 0, len(grid))

	var count int

	for x, row := range grid {

		n := make([]string, 0, len(row))

		for y, val := range row {
			if val == "." {
				fmt.Print(".")
				n = append(n, ".")
				continue
			}

			if getAdjacentCount(grid, x, y) < 4 {
				fmt.Print("x")
				n = append(n, ".")
				count++
			} else {
				fmt.Print("@")
				n = append(n, "@")
			}
		}

		fmt.Println("")
		out = append(out, n)
	}

	return out, count
}

func getAdjacentCount(grid [][]string, row int, column int) int {
	if !exists(grid, row, column) {
		return 0
	}

	if grid[row][column] == "." {
		return 0
	}

	var count int

	for _, ch := range checks {
		if hasRoll(grid, row+ch.X, column+ch.Y) {
			count++
		}
	}

	return count
}

var checks = []Position{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type Position struct {
	X int
	Y int
}

func exists(grid [][]string, row int, column int) bool {
	if row < 0 || column < 0 {
		return false
	}

	if row >= len(grid) {
		return false
	}

	if column >= len(grid[row]) {
		return false
	}
	return true
}

func hasRoll(grid [][]string, row int, column int) bool {
	if !exists(grid, row, column) {
		return false
	}

	value := grid[row][column]

	return value == "@"
}
