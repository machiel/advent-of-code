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
	ops := parseOperations(transform(parseInput(readInput())))

	var total int

	for _, op := range ops {
		total += op.Solve()
	}

	fmt.Println(total)
}

func parseInput(in iter.Seq[string]) [][]string {
	sizes := map[int]int{}
	lines := make([]string, 0)

	for line := range in {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)

		for i, field := range fields {
			if _, ok := sizes[i]; !ok {
				sizes[i] = len(field)
				continue
			}

			if len(field) > sizes[i] {
				sizes[i] = len(field)
			}
		}

		lines = append(lines, line)
	}

	out := make([][]string, 0)

	for _, line := range lines {
		fields := make([]string, len(sizes))

		for i := 0; i < len(sizes); i++ {
			value := line[:sizes[i]]
			fields[i] = value

			if i < len(sizes)-1 {
				line = line[sizes[i]+1:]
			}
		}

		out = append(out, fields)
	}

	return out
}

func readInput() iter.Seq[string] {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return strings.SplitSeq(string(input), "\n")
}

func parse(in string) int {
	v, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return v
}

func transform(in [][]string) [][]string {
	out := make([][]string, len(in[0]))

	for _, row := range in {
		for i, value := range row {
			out[i] = append(out[i], value)
		}
	}

	return out
}

type Type string

const (
	Add      Type = "+"
	Multiply Type = "*"
)

type Operation struct {
	Operation Type
	Values    []int
}

func (o Operation) Solve() int {
	var result int

	for i, v := range o.Values {
		switch o.Operation {
		case Add:
			result = result + v
		case Multiply:
			if i == 0 {
				result = v
			} else {
				result = result * v
			}
		default:
			panic("Unknown operation: " + o.Operation)
		}
	}

	return result
}

func resolveValues(in []string) []int {
	numbers := make([]string, len(in[0]))

	l := len(in[0]) - 1

	for i := l; i >= 0; i-- {
		for _, v := range in {
			numbers[l-i] = numbers[l-i] + string(v[i])
		}
	}

	out := make([]int, 0, len(numbers))

	for _, n := range numbers {
		n = strings.TrimSpace(n)
		out = append(out, parse(n))
	}

	return out
}

func parseOperation(in []string) Operation {
	return Operation{
		Operation: parseType(in[len(in)-1]),
		Values:    resolveValues(in[:len(in)-1]),
	}
}

func parseOperations(in [][]string) []Operation {
	out := make([]Operation, 0, len(in))

	for _, row := range in {
		out = append(out, parseOperation(row))
	}

	return out
}

func parseType(in string) Type {
	switch strings.TrimSpace(in) {
	case "+":
		return Add
	case "*":
		return Multiply
	}

	panic("Input " + in + "not known")
}
