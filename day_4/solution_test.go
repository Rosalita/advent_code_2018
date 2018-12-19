package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitIntoShifts(t *testing.T) {

	input := []string{
		"[1518-02-28 23:56] Guard #971 begins shift",
		"[1518-03-01 00:15] falls asleep",
		"[1518-03-01 00:54] wakes up",
		"[1518-03-02 00:02] Guard #3079 begins shift",
		"[1518-03-02 00:08] falls asleep",
		"[1518-03-02 00:48] wakes up",
		"[1518-03-02 00:51] falls asleep",
		"[1518-03-02 00:52] wakes up",
		"[1518-03-02 23:58] Guard #653 begins shift",
		"[1518-03-03 00:35] falls asleep",
		"[1518-03-03 00:55] wakes up",
	}

	line1 := []string{
		"#971", "asleep15", "wakes54",
	}

	line2 := []string{
		"#3079", "asleep08", "wakes48", "asleep51", "wakes52",
	}

	line3 := []string{
		"#653", "asleep35", "wakes55",
	}

	result := [][]string{
		line1, line2, line3,
	}

	tests := []struct {
		input  []string
		result [][]string
	}{
		{input, result},
	}
	for _, test := range tests {
		result := splitIntoShifts(test.input)
		assert.Equal(t, test.result, result)
	}

}
