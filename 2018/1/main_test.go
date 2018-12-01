package main

import (
	"testing"
)

func Test_sumInput(t *testing.T) {
	testInputs := []struct {
		input    []int
		expected int
	}{
		{[]int{1, -1}, 0},
		{[]int{1, 2, 3, 4, 6}, 16},
		{[]int{-1, -2, -3, -4, -6}, -16},
	}

	for i, test := range testInputs {
		actualOutput := sumInput(test.input)
		if actualOutput != test.expected {
			t.Errorf("(%v) testInputs(%v) expected: %v actual: %v", i, test.input, test.expected, actualOutput)
		}
	}
}
func Test_DuplicateFrequency(t *testing.T) {
	testInputs := []struct {
		input  []int
		answer int
	}{
		{[]int{+1, -1}, 0},
		{[]int{+1, +2, -2}, 1},
		{[]int{+3, +3, +4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
	}

	for i, test := range testInputs {
		actualOutput, err := duplicateFrequency(test.input)
		if err != nil {
			t.Errorf("duplicateFrequency returned an error, but it shouldn't, %v", err)
		}
		if actualOutput != test.answer {
			t.Errorf("(%v) duplicateFrequency(%v) fail: expected: %v, actual: %v", i, test.input, test.answer, actualOutput)
		}
	}
}
