package bleach_test

import (
	"testing"

	"github.com/Kangaroux/go-bleach"
	"github.com/stretchr/testify/require"
)

func TestToString(t *testing.T) {
	type testPair struct {
		input    interface{}
		expected string
	}

	testInputs := []testPair{
		{0, "0"},
		{123, "123"},
		{3.14, "3.14"},
		{-3.14, "-3.14"},
		{true, "true"},
		{false, "false"},
		{"foo bar", "foo bar"},
		{nil, "null"},
	}

	m := bleach.ToString()

	for _, input := range testInputs {
		require.Equal(t, input.expected, m.Mutate(input.input))
	}
}

func TestToInt(t *testing.T) {
	type testPair struct {
		input    interface{}
		expected int64
	}

	testInputs := []testPair{
		{0, 0},
		{-1, -1},
		{3.14, 3},
		{3.99, 3},
		{-3.99, -3},
		{uint(5), 5},
		{true, 1},
		{false, 0},
		{"1", 1},
		{"3.14", 3},
		{"-3.99", -3},
		{nil, 0},
		{"", 0},
		{"foo", 0},
		{"foo 123", 0},
	}

	m := bleach.ToInt()

	for _, input := range testInputs {
		require.Equal(t, input.expected, m.Mutate(input.input), "input (%T): %#v", input.input, input.input)
	}
}
