package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRepetitive(t *testing.T) {
	tests := map[string]bool{
		"0":         false,
		"1":         false,
		"2":         false,
		"10":        false,
		"11":        true,
		"15":        false,
		"20":        false,
		"22":        true,
		"100":       false,
		"101":       false,
		"111":       true,
		"1000":      false,
		"1010":      true,
		"121210":    false,
		"121212":    true,
		"88888888":  true,
		"12341234":  true,
		"303303303": true,
		"20202020":  true,
	}

	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			result := FindRepetive(input)

			assert.Equal(t, expected, result)
		})
	}
}
