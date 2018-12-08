package main

import (
	"strings"
	"testing"
)

func Test_graphTest(t *testing.T) {
	inputTests := []struct {
		input          []string
		expectedOutput string
	}{
		{[]string{
			"Step C must be finished before step A can begin.",
			"Step C must be finished before step F can begin.",
			"Step A must be finished before step B can begin.",
			"Step A must be finished before step D can begin.",
			"Step B must be finished before step E can begin.",
			"Step D must be finished before step E can begin.",
			"Step F must be finished before step E can begin.",
		},
			"CABDFE",
		},
		{[]string{
			"Step J must be finished before step C can begin.",
			"Step I must be finished before step J can begin.",
			"Step J must be finished before step A can begin.",
			"Step A must be finished before step Q can begin.",
			"Step X must be finished before step I can begin.",
		}, "XIJACQ",
		},
	}
	for _, test := range inputTests {
		parsedLines := parseLines(test.input)
		graph, dependentGraph := createConnections(parsedLines)
		rootNodes := findRootNodes(graph)
		path := graphTraverse(graph, dependentGraph, rootNodes)
		if strings.Join(path, "") != test.expectedOutput {
			t.Errorf("Expected path: %v, actual: %v", test.expectedOutput, path)
		}
	}

}
