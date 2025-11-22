package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " you'RE  kilLinG mE  sMALLs  ",
			expected: []string{"you're", "killing", "me", "smalls"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Fatalf("input: %v yields %v; test failed", c.input, actual)

		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if !reflect.DeepEqual(expectedWord, word) {
				t.Fatalf("expected word: %v, actual word: %v", expectedWord, word)
			}

		}

	}

}
