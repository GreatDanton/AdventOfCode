package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	initialState, growingRules := readInput()
	initSize := len(initialState)
	thirdEmpty := make([]string, initSize)
	for i := 0; i < len(thirdEmpty); i++ {
		thirdEmpty[i] = "."
	}
	indexMovedFor := len(thirdEmpty)

	// we have initial state in the middle
	allPots := append(thirdEmpty, strings.Split(initialState, "")...)
	allPots = append(allPots, thirdEmpty...)

	for gen := 0; gen < 20; gen++ {
		allPots = createGeneration(allPots, growingRules)
	}
	lastGenScore := calculatePotScore(allPots, indexMovedFor)
	fmt.Println("Part 1, 20th generation pot score:", lastGenScore)

	// part 2
	hugeEmpty := make([]string, initSize*10)
	for i := 0; i < len(hugeEmpty); i++ {
		hugeEmpty[i] = "."
	}
	indexMovedFor = len(hugeEmpty)
	part2Pots := append(hugeEmpty, strings.Split(initialState, "")...)
	part2Pots = append(part2Pots, hugeEmpty...)
	previous := 0
	difference := 0
	// 200 is magic number to get the constant difference for each generation iteration
	const reachConstantIncrease = 200
	for gen := 0; gen < reachConstantIncrease; gen++ {
		part2Pots = createGeneration(part2Pots, growingRules)
		score := calculatePotScore(part2Pots, indexMovedFor)
		difference = score - previous
		previous = score
	}
	currentScore := calculatePotScore(part2Pots, indexMovedFor)
	const endGeneration = 50000000000
	fmt.Printf("Part2 pot score after %d years: %d\n", endGeneration, currentScore+(endGeneration-reachConstantIncrease)*difference)

}

func createGeneration(allPots []string, growingRules [][]string) []string {
	// 2 is magic number we can't apply any rules to left/right most elements
	nextGeneration := append([]string{}, allPots...)
	for pot := 2; pot < len(allPots)-2; pot++ {
		for r := 0; r < len(growingRules); r++ {
			rule := growingRules[r]
			r1, r2, middle, r3, r4 := rule[0], rule[1], rule[2], rule[3], rule[4]
			p1, p2, middlePot, p3, p4 := allPots[pot-2], allPots[pot-1], allPots[pot], allPots[pot+1], allPots[pot+2]
			// check if middle pot matches with middle rule element
			if middlePot == middle {
				// we have a match, next generation this position will have a pot
				if r1 == p1 && r2 == p2 && r3 == p3 && r4 == p4 {
					nextGeneration[pot] = "#"
					break
				}
			}
			// check last element did not break (change an element to the pot)
			if r == len(growingRules)-1 {
				nextGeneration[pot] = "."
			}
		}
	}
	return nextGeneration
}

func calculatePotScore(pots []string, indexMovedFor int) int {
	score := 0
	for i := 0; i < len(pots); i++ {
		if pots[i] == "#" {
			score += i - indexMovedFor
		}
	}
	return score
}

func readInput() (string, [][]string) {
	file, err := os.Open("input.txt")
	// file, err := os.Open("testInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	line := 0
	initialState := ""
	rules := [][]string{}
	for scanner.Scan() {
		if line == 0 {
			initialState = strings.Trim(strings.Split(scanner.Text(), ":")[1], " ")
		} else if line == 1 {
			line++
			continue
		} else {
			rule := scanner.Text()
			if strings.Contains(rule, "=> #") {
				row := strings.Split(strings.Split(rule, " ")[0], "")
				rules = append(rules, row)
			}
		}
		line++
	}
	return initialState, rules
}
