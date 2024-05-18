package main

import "testing"

// takes a pointer to 'T' for tests in the testing package
// you tell golang when a test has failed explicitly by calling T.error, rather than
// goland evaluating an assertion like in vitest for JS... I think...
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		// Check length of input
		if len(actual) != len(cs.expected) {
			t.Errorf("Lengths not equal!\n ACTUAL: %v\nEXPECTED: %v",
				len(actual),
				len(cs.expected),
			)
			continue
		}
		// Check contents of string
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v",
					actualWord,
					expectedWord,
				)
			}
		}

	}

}
