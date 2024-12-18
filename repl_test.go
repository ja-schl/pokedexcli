package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hEllo World   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "charizard Bulbasaur PICKACHU",
			expected: []string{"charizard", "bulbasaur", "pickachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual output, %d, does not match expected length of %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("%q does not match expected word %q", word, expected)
			}
		}
	}
}
