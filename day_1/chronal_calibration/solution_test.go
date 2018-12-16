package main

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestApplyChanges(t *testing.T) {

	tests := []struct {
		input  []string
		result int
	}{
		{[]string{"+1", "+1", "+1"}, 3},
		{[]string{"+1", "+1", "-2"}, 0},
		{[]string{"-1", "-2", "-3"}, -6},
	}
	for _, test := range tests {
		result := applyChanges(test.input)
		assert.Equal(t, test.result, result)
	}

}

func TestCheckResult(t *testing.T){
	tests := []struct {
		result  int
		history []int
		isInHistory bool
	}{
		{4, []int{1,2,3}, false},
		{0, []int{3,2,0}, true},
	}
	for _, test := range tests {
		isInHistory := checkResult(test.result, test.history)
		assert.Equal(t, test.isInHistory, isInHistory)
	}
}
