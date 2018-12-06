package main

import (
	"fmt"
	"testing"
)

func Test_claimVirtualPoints(t *testing.T) {
	inputTests := []struct {
		source         point
		r              int
		expectedPoints []point
	}{
		{source: point{x: 2, y: 2}, r: 1, expectedPoints: []point{
			{x: 2, y: 3},
			{x: 3, y: 2},
			{x: 2, y: 1},
			{x: 1, y: 2},
		},
		},

		{source: point{x: 2, y: 2}, r: 2, expectedPoints: []point{
			{x: 2, y: 4},
			{x: 4, y: 2},
			{x: 2, y: 0},
			{x: 0, y: 2},
			//
			{x: 1, y: 3},
			{x: 3, y: 3},
			{x: 3, y: 1},
			{x: 1, y: 1},
		},
		},
	}
	for _, test := range inputTests {
		points := claimVirtualPoints(test.source, test.r)
		fmt.Println(points)
		if len(points) != len(test.expectedPoints) {
			t.Errorf("Expected: %v, actual: %v", test.expectedPoints, points)
		}
		for _, p := range points {
			contain := contains(test.expectedPoints, p)
			if !contain {
				t.Errorf("Point %v is not contained in expected array: %v", p, test.expectedPoints)
			}

		}
	}

}
