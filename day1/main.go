package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()

	instructions := make([]Rotation, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}

		instruction, err := parseInstruction(line)
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, instruction)
	}

	state := State(50)
	var count int

	fmt.Println("Starting at state", state)

	for _, instruction := range instructions {
		var seenZeros int
		state, seenZeros = state.Apply(instruction)

		count += seenZeros

		if state == State(0) {
			count++
		}

		if seenZeros > 0 {
			fmt.Println("Applying instruction", instruction, "resulted in", state, "pointing", seenZeros, "times at 0")
		} else {
			fmt.Println("Applying instruction", instruction, "resulted in", state)
		}
	}

	fmt.Println("We have seen zero", count, "times")
}

func readInput() []string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return slices.Collect(strings.SplitSeq(string(input), "\n"))
}

func parseInstruction(in string) (Rotation, error) {
	if len(in) <= 1 {
		return Rotation{}, fmt.Errorf("input too short: %s", in)
	}

	direction, err := parseDirection(in[0])
	if err != nil {
		return Rotation{}, err
	}

	offset, err := strconv.Atoi(in[1:])
	if err != nil {
		return Rotation{}, err
	}

	return Rotation{
		Direction: direction,
		Offset:    offset,
	}, nil
}

func parseDirection(in byte) (Direction, error) {
	switch in {
	case 'L':
		return DirectionLeft, nil
	case 'R':
		return DirectionRight, nil
	default:
		return Direction(""), errors.New("invalid directional input")
	}
}

type Direction string

var (
	DirectionLeft  Direction = "LEFT"
	DirectionRight Direction = "RIGHT"
)

type Rotation struct {
	Direction Direction
	Offset    int
}

type State int

func (s State) Apply(rot Rotation) (State, int) {
	if rot.Direction == DirectionLeft {
		return s.decrease(rot.Offset)
	}

	if rot.Direction == DirectionRight {
		return s.increase(rot.Offset)
	}

	return s, 0
}

func (s State) decrease(offset int) (State, int) {
	if s == 0 {
		s = 100
	}

	roundtrips := (offset / 100)
	s = (State((roundtrips+1)*100) + s) - State(offset)

	if s >= 100 {
		return s % 100, roundtrips
	}

	return s, roundtrips + 1
}

func (s State) increase(offset int) (State, int) {
	newState := s + State(offset)
	seenZeros := int(newState) / 100

	newState = newState % 100

	if newState == 0 {
		seenZeros = seenZeros - 1
	}

	return newState, seenZeros
}

func Rot(d Direction, o int) Rotation {
	return Rotation{
		Direction: d,
		Offset:    o,
	}
}
