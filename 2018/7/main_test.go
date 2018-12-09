package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_graphTest(t *testing.T) {
	inputTests := []struct {
		input          []string
		expectedOutput string
		maxWorkingTime int
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
			"CABDFE", 898,
		},
		{[]string{
			"Step J must be finished before step C can begin.",
			"Step I must be finished before step J can begin.",
			"Step J must be finished before step A can begin.",
			"Step A must be finished before step Q can begin.",
			"Step X must be finished before step I can begin.",
		}, "XIJACQ", 0,
		},
	}
	for _, test := range inputTests {
		parsedLines := parseLines(test.input)
		graph, dependencyGraph := createConnections(parsedLines)
		rootNodes := findRootNodes(graph, dependencyGraph)
		fmt.Println("Graph:", graph, "Dep graph: ", dependencyGraph)
		path := graphTraverse(graph, dependencyGraph, rootNodes)
		if strings.Join(path, "") != test.expectedOutput {
			t.Errorf("Expected path: %v, actual: %v", test.expectedOutput, path)
		}

		/* 		graph, dependencyGraph = createConnections(parsedLines)
		   		time := calculateTime(path, graph, dependencyGraph)
		   		if time != test.maxWorkingTime {
		   			t.Errorf("Time expected: %d, time actual: %d\n", test.maxWorkingTime, time)
		   		} */
	}

}
