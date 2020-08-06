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
