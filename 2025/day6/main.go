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
	input := make([][]string, 0)

	for line := range readInput() {
		if line == "" {
			continue
		}

		input = append(input, strings.Fields(line))
	}

	ops := transform(input)

	var total int

	for _, op := range ops {
		total += op.Solve()
	}

	fmt.Println(total)
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

func transform(in [][]string) []Operation {
	out := make([]Operation, len(in[0]))

	for lineNo, row := range in {
		for index, v := range row {
			if lineNo == len(in)-1 {
				switch v {
				case "+":
					out[index].Operation = Add
				case "*":
					out[index].Operation = Multiply
				default:
					panic("Input " + v + "not known")
				}
			} else {
				out[index].Values = append(out[index].Values, parse(v))
			}
		}
	}

	return out
}
