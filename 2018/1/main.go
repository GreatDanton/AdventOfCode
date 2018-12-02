package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 1. day: https://adventofcode.com/2018/day/1
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := readInput(file)
	if err != nil {
		log.Fatal(err)
	}

	sum := sumInput(input)
	fmt.Printf("Sum of the numbers is: %v\n", sum)

	duplicateFrequency, err := duplicateFrequency(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Duplicate frequency is: %d\n", duplicateFrequency)
}

// read file and prepare an array of inputs for further processing
func readInput(file *os.File) ([]int, error) {
	scanner := bufio.NewScanner(file)
	input := []int{}
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error occurred while processing a line: %v", err)
			continue
		}
		input = append(input, num)
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return input, nil
}

// 1. part
// ---------------------------------------------------------
// sum all numbers from the input
func sumInput(input []int) int {
	sum := 0
	for _, num := range input {
		sum += num
	}
	return sum
}

// 2. part
// ---------------------------------------------------------
// find the duplicate frequency that appears in the input. If the frequency
// is not found on the first try, it will iterate over the same input once
// again (on weird input such as [1,2,3] the loop will never finish)
func duplicateFrequency(input []int) (int, error) {
	frequency := 0
	frequencies := map[int]int{frequency: 1}
	for i := 0; i < len(input); {
		frequency += input[i]
		_, exists := frequencies[frequency]
		if exists {
			return frequency, nil
		}
		frequencies[frequency]++

		// if we are on the last element and we still did not find the
		// duplicated frequency iterate over the same loop once again
		i++
		if i == len(input) {
			i = 0
		}
	}

	// this will never happen, but we can't compile without this statement
	return 0, fmt.Errorf("There is no frequency that would be duplicated")
}
