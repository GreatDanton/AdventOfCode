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
	graph, dependencyGraph := createConnections(parsedLines)
	rootNodes := findRootNodes(graph, dependencyGraph)
	path := graphTraverse(graph, dependencyGraph, rootNodes)
	fmt.Println("Path is: ", strings.Join(path, ""))

	// part 2
	// create both graphs once again, as the part 1 is modifying them
	graph, dependencyGraph = createConnections(parsedLines)
	time := calculateTime(path, dependencyGraph)
	fmt.Printf("Time it takes for %d workers: %d\n", numOfWorkers, time)
}

func graphTraverse(graph map[string][]string, dependencyGraph map[string][]string, rootNodes []string) []string {
	path := []string{}
	notVisited := []string{}
	notVisited = append(notVisited, rootNodes...)
	rootNodeTraverse(graph, dependencyGraph, "", &path, notVisited)
	return path
}

func rootNodeTraverse(graph map[string][]string, dependencyGraph map[string][]string, rootNode string, path *[]string, notVisited []string) {
	if rootNode != "" {
		*path = append(*path, rootNode)
		connections := graph[rootNode]
		// filter out connections that do not have dependencies sorted out
		for _, connection := range connections {
			if checkDependencies(dependencyGraph, connection, *path) && !contains(notVisited, connection) {
				notVisited = append(notVisited, connection)
			}
		}
	}

	if len(notVisited) > 0 {
		sortNodesAZ(notVisited)
		nextElement := notVisited[0]
		notVisited = notVisited[1:]
		rootNodeTraverse(graph, dependencyGraph, nextElement, path, notVisited)
	}
}

func checkDependencies(dependencyGraph map[string][]string, connection string, path []string) bool {
	nodesFrom := dependencyGraph[connection]
	newNodesFrom := make([]string, 0, len(nodesFrom))
	for _, node := range nodesFrom {
		if !contains(path, node) {
			newNodesFrom = append(newNodesFrom, node)
		}
	}
	if len(newNodesFrom) == 0 {
		return true
	}
	dependencyGraph[connection] = newNodesFrom
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
func findRootNodes(graph map[string][]string, dependencyGraph map[string][]string) []string {
	connections := map[string]int{}
	for node, nodeConn := range graph {
		connections[node]++
		for _, n := range nodeConn {
			connections[n]++
		}
	}

	rootNodes := []string{}
	for node, count := range connections {
		_, exist := dependencyGraph[node]
		if count == 1 && !exist {
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

// part 2

// LETTERS of the alphabet
const LETTERS string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// TIME for each letter
const TIME int = 60

func getTimeForLetter(letter string) int {
	// A starts with 1
	t := strings.Index(LETTERS, letter) + 1
	if t > 0 {
		return TIME + t
	}
	// wrong input
	return -1
}

const numOfWorkers = 5

var workers = make([]int, numOfWorkers)
var workerTasks = make([]string, numOfWorkers) // working on which part of path?

func calculateTime(path []string, dependenciesGraph map[string][]string) int {
	// check for path
	time := 0
	var donePath = []string{}
	for {
		time++
		// pick a task and assign it to worker
		pickPossibleTask(path, &donePath, dependenciesGraph)

		// decrease workers time in each iteration
		for index := range workers {
			if workers[index] > 0 {
				workers[index]--
			}
		}

		// check if all workers are done and completed path == initial path
		stillWorking := false
		for index, workingTime := range workers {
			if workingTime > 0 {
				stillWorking = true
			} else {
				// worker is idle or just completed the work
				if workerTasks[index] != "" { // just completed the work
					// add path to array of done paths, and clear current tasks
					donePath = append(donePath, workerTasks[index])
					workerTasks[index] = ""
					// try to pick the next task for any worker
					pickingTask := pickPossibleTask(path, &donePath, dependenciesGraph)
					// if any worker set still working flag to true, skip this part
					// if there is nobody working, one worker might just complete its task
					// and was assigned a new task
					if !stillWorking {
						stillWorking = pickingTask
					}
				}
			}
		}
		if !stillWorking {
			break
		}
	}
	return time
}

// will try to pick all possible tasks?
func pickPossibleTask(finalPath []string, donePath *[]string, dependenciesGraph map[string][]string) bool {
	remainingPath := calculateRemainingPath(finalPath, donePath)
	for _, node := range remainingPath {
		possible := checkDependencies(dependenciesGraph, node, *donePath)
		if possible {
			return assignWorker(node) // assign as much workers as possible
		}
	}
	return false
}

func assignWorker(letter string) bool {
	for index, time := range workers {
		if time == 0 { // idle worker
			workingTime := getTimeForLetter(letter)
			if workingTime < 0 { // handler error
				fmt.Println("There was an error with parsing letter")
				return false
			}
			// assigned worker
			workers[index] = workingTime
			workerTasks[index] = letter
			return true
		}
	}
	return false
}

func calculateRemainingPath(finalPath []string, donePath *[]string) []string {
	remaining := []string{}
	for _, currentPath := range finalPath {
		if !contains(*donePath, currentPath) && !contains(workerTasks, currentPath) {
			remaining = append(remaining, currentPath)
		}
	}
	return remaining
}
