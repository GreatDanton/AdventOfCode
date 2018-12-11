package main

import (
	"testing"
)

func Test_parseTest(t *testing.T) {
	inputTests := []struct {
		input       []int
		expectedSum int
	}{
		{[]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}, 138},
	}

	for _, test := range inputTests {
		//actualSum := sumMetadata(test.input)
		actualSum, _ := parse(test.input)

		if test.expectedSum != actualSum {
			t.Errorf("Expected sum: %v, actual sum: %v", test.expectedSum, actualSum)
		}
	}
}
