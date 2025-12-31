package cmd

import "testing"

func TestCleanInput( t * testing.T) {
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
			input: "Hello World",
			expected:  []string{
				"hello", "world",
			},
		},
	}

	for  _, csc := range cases {
		real := CleanInput(csc.input)
		if len(real) != len(csc.expected){
			t.Errorf("The lengths are not equal as expected %v, %v", real, csc.expected)
			continue
		}

		for i := range real {
			realWord := real[i]
			expectedWord := csc.expected[i]

			if realWord != expectedWord {
				t.Errorf("The Word are not equal as expected %v, %v", realWord, expectedWord)
			}
		}

	}
}