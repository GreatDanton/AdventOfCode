package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		lines = append(lines, strings.Split(scanner.Text(), " ")...)
	}

	nums := convertToNum(lines)
	result, _ := parse(nums)
	fmt.Println(result)

	/* nums := convertToNum(lines)
	   //sumMetadata(nums) */
}

func parse(input []int) (int, []int) {
	header := input[:2]
	input = input[2:]
	children := header[0]
	metadata := header[1]
	allTotal := 0

	// recursive parse for each children until we get to children 0
	for i := 0; i < children; i++ {
		total := 0
		total, input = parse(input)
		allTotal += total
	}

	if children == 0 {
		allTotal += Sum(input[:metadata])
		return allTotal, input[metadata:]
	}
	// TODO: finish the else block

	allTotal += Sum(input[:metadata])
	return allTotal, input[metadata:]
}

// Sum calculates the sum of all provided elements in the slice
func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func convertToNum(numArr []string) []int {
	ints := make([]int, 0, len(numArr))
	for _, line := range numArr {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, num)
	}
	return ints
}
