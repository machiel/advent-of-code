package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_Apply(t *testing.T) {
	tests := map[string]struct {
		State    State
		Input    Rotation
		Expected State
	}{
		"82 with L30": {
			State:    State(82),
			Input:    Rot(DirectionLeft, 30),
			Expected: State(52),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, _ := test.State.Apply(test.Input)

			assert.Equal(t, test.Expected, out)
		})
	}
}

func TestState_increase(t *testing.T) {
	tests := map[string]struct {
		State       State
		Input       int
		OutputState State
		SeenZeros   int
	}{
		"50+100": {
			State:       State(50),
			Input:       100,
			OutputState: State(50),
			SeenZeros:   1,
		},
		"50+5": {
			State:       State(50),
			Input:       5,
			OutputState: State(55),
			SeenZeros:   0,
		},
		"50+50": {
			State:       State(50),
			Input:       50,
			OutputState: State(0),
			SeenZeros:   0,
		},
		"50+51": {
			State:       State(50),
			Input:       51,
			OutputState: State(1),
			SeenZeros:   1,
		},
		"0+101": {
			State:       State(0),
			Input:       101,
			OutputState: State(1),
			SeenZeros:   1,
		},
		"0+99": {
			State:       State(0),
			Input:       99,
			OutputState: State(99),
			SeenZeros:   0,
		},
		"0+990": {
			State:       State(0),
			Input:       990,
			OutputState: State(90),
			SeenZeros:   9,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, seenZeros := test.State.increase(test.Input)

			assert.Equal(t, test.OutputState, out)
			assert.Equal(t, test.SeenZeros, seenZeros)
		})
	}
}

func TestState_decrease(t *testing.T) {
	tests := map[string]struct {
		State       State
		Input       int
		OutputState State
		SeenZeros   int
	}{
		"50-100": {
			State:       State(50),
			Input:       100,
			OutputState: State(50),
			SeenZeros:   1,
		},
		"50-5": {
			State:       State(50),
			Input:       5,
			OutputState: State(45),
			SeenZeros:   0,
		},
		"50-50": {
			State:       State(50),
			Input:       50,
			OutputState: State(0),
			SeenZeros:   0,
		},
		"50-51": {
			State:       State(50),
			Input:       51,
			OutputState: State(99),
			SeenZeros:   1,
		},
		"0-101": {
			State:       State(0),
			Input:       101,
			OutputState: State(99),
			SeenZeros:   1,
		},
		"0-99": {
			State:       State(0),
			Input:       99,
			OutputState: State(1),
			SeenZeros:   0,
		},
		"0-990": {
			State:       State(0),
			Input:       990,
			OutputState: State(10),
			SeenZeros:   9,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, seenZeros := test.State.decrease(test.Input)

			assert.Equal(t, test.OutputState, out)
			assert.Equal(t, test.SeenZeros, seenZeros)
		})
	}
}
