package main

import(
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestHasNumberOfSameCharacters(t *testing.T) {

	tests := []struct {
		input  string
		number int
		expected bool
	}{
		{"abcdef", 2, false},
		{"abcdef", 3, false},
		{"bababc", 2, true},
		{"bababc", 3, true},
		{"abbcde", 2, true},
		{"abbcde", 3, false},
		{"abcccd", 3, true},
		{"abcccd", 2, false},
		{"aabcdd", 2, true},
		{"abcdee", 2, true},
		{"ababab", 3, true},

	}
	for _, test := range tests {
		result := hasNumberOfSameCharacters(test.input, test.number)
		assert.Equal(t, test.expected, result)
	}

}
