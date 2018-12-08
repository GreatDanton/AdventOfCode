package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		// TODO: handle error
	}
	parsedLines := parseLines(lines)
	graph, dependentGraph := createConnections(parsedLines)
	rootNodes := findRootNodes(graph)
	path := graphTraverse(graph, dependentGraph, rootNodes)
	fmt.Println("Path is: ", strings.Join(path, ""))
}

func graphTraverse(graph map[string][]string, dependentGraph map[string][]string, rootNodes []string) []string {
	path := []string{}
	notVisited := []string{}
	notVisited = append(notVisited, rootNodes...)
	rootNodeTraverse(graph, dependentGraph, "", &path, notVisited)
	return path
}

func rootNodeTraverse(graph map[string][]string, dependentGraph map[string][]string, rootNode string, path *[]string, notVisited []string) {
	if rootNode != "" {
		*path = append(*path, rootNode)
		connections := graph[rootNode]
		// filter out connections that do not have dependencies sorted out
		for _, connection := range connections {
			if checkDependencies(dependentGraph, connection, *path) && !contains(notVisited, connection) {
				notVisited = append(notVisited, connection)
			}
		}
	}

	if len(notVisited) > 0 {
		sortNodesAZ(notVisited)
		nextElement := notVisited[0]
		notVisited = notVisited[1:]
		rootNodeTraverse(graph, dependentGraph, nextElement, path, notVisited)
	}
}

func checkDependencies(dependentGraph map[string][]string, connection string, path []string) bool {
	nodesFrom := dependentGraph[connection]
	newNodesFrom := make([]string, 0, len(nodesFrom))
	for _, node := range nodesFrom {
		if !contains(path, node) {
			newNodesFrom = append(newNodesFrom, node)
		}
	}
	if len(newNodesFrom) == 0 {
		return true
	}
	dependentGraph[connection] = newNodesFrom
	return false
}

func createConnections(parsedLines [][]string) (map[string][]string, map[string][]string) {
	connections := map[string][]string{}
	dependencies := map[string][]string{}
	for _, line := range parsedLines {
		from := line[0]
		to := line[1]
		connections[from] = append(connections[from], to)
		dependencies[to] = append(dependencies[to], from)
	}
	return connections, dependencies
}

// find nodes that nobody is pointing to
func findRootNodes(graph map[string][]string) []string {
	connections := map[string]int{}
	for node, nodeConn := range graph {
		connections[node]++
		for _, n := range nodeConn {
			connections[n]++
		}
	}

	rootNodes := []string{}
	for node, count := range connections {
		if count == 1 {
			rootNodes = append(rootNodes, node)
		}
	}
	return rootNodes
}

func parseLines(input []string) [][]string {
	steps := make([][]string, 0, len(input))
	for _, line := range input {
		before := ""
		after := ""
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &before, &after)
		steps = append(steps, []string{before, after})
	}
	return steps
}

func contains(arr []string, element string) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}
	return false
}

func sortNodesAZ(nodes []string) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i] < nodes[j]
	})
}
