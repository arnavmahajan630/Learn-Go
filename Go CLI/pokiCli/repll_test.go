package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input string
		expected []string
	} {
		{
			input: "hello world",
			expected: []string{
				"hello", "world",
			},
		},

		{
			input: "Hello WORLD",
			expected: []string{
				"hello", "world",
			},
		},
	}

	for _, cs := range cases {
		real := cleanInput(cs.input)
		if len(real) != len(cs.expected){
			t.Errorf("The lengths are not equal: %v vs %v", len(real), len(cs.expected))
			continue
		}

		for i := range real {
			realWord := real[i]
			expectedWord := cs.expected[i]

			if realWord != expectedWord {
				t.Errorf("The words are not equal: %v vs %v", realWord, expectedWord)
				continue
			}
		}
	}
}
