package main

import "testing"


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{	
		{
			input: "  hello world   ",
			expected: []string{"hello","world"},
		},
		{
			input: " my oh my!",
			expected: []string{"my","oh","my!"},
		},
		{
			input: " Go is awesome   ",
			expected: []string{"go","is","awesome"},
		},
		{
			input: " here is tougher case, yeah, tough, like your dad with bottle of whiskey in his bElly   ",
			expected: []string{"here", "is", "tougher", "case", "yeah", "tough", "like", "your", "dad", "with", "bottle", "of", "whiskey", "in", "his", "belly"},
		},
		// add more cases
	}
	
	for _, c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The expected length of slice doest match expected\nExpected: %d\nActual: %d",len(c.expected),len(c.input))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word in a slice doesnt match\nEXPECTED: %s\nACTUAL: %s", expectedWord, word)
			}
		}
	}
}