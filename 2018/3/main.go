package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type elfFabric struct {
	ID       int
	position []int
	size     []int
}

type point struct {
	x int
	y int
}

// 3.day: https://adventofcode.com/2018/day/3
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// part 1
	fabricArr := readInput(file)
	count := calculateClaimedFabric(fabricArr)
	fmt.Printf("Square inches taken by multiple elfs: %d\n", count)

	// part 2
	fabricID := getLoneFabric(fabricArr)
	fmt.Printf("Lone fabric id: %v\n", fabricID)
}

// part 1
// ----------------------------------------
func readInput(file *os.File) []elfFabric {
	scanner := bufio.NewScanner(file)
	fabricArr := []elfFabric{}
	r, err := regexp.Compile(`\d+`)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parsedNumbers := r.FindAllString(line, -1)
		if parsedNumbers == nil || len(parsedNumbers) < 5 {
			// regex did not find any numbers or less than 5, malformed line
			continue
		}

		intsArr := arrayToInt(parsedNumbers)
		id := intsArr[0]
		pos := []int{intsArr[1], intsArr[2]}
		size := []int{intsArr[3], intsArr[4]}
		fabricArr = append(fabricArr, elfFabric{ID: id, position: pos, size: size})
	}
	return fabricArr
}

func arrayToInt(arr []string) []int {
	out := make([]int, 0, len(arr))
	for _, element := range arr {
		parsedInt, err := strconv.Atoi(element)
		if err != nil {
			// TODO: handle error, this line is malformed
			fmt.Println(err)
		}
		out = append(out, parsedInt)
	}
	return out
}

func createFabricClaims(fabricArr []elfFabric) map[point]int {
	takenFabric := map[point]int{}
	// holy shit this has awful performance
	for _, fabric := range fabricArr {
		for addX := 1; addX <= fabric.size[0]; addX++ {
			for addY := 1; addY <= fabric.size[1]; addY++ {
				point := point{x: fabric.position[0] + addX, y: fabric.position[1] + addY}
				takenFabric[point]++
			}
		}
	}
	return takenFabric
}

func calculateClaimedFabric(fabricArr []elfFabric) int {
	takenFabric := createFabricClaims(fabricArr)
	count := 0
	for _, v := range takenFabric {
		// count inches taken by more than two elfs
		if v >= 2 {
			count++
		}
	}
	return count
}

// part 2
// --------------------------------
// check every point of the claimed fabric, if one tiny part of it overlaps with another
// fabric return an invalid id
func checkPoints(fabric *elfFabric, takenFabric map[point]int) int {
	for addX := 1; addX <= fabric.size[0]; addX++ {
		for addY := 1; addY <= fabric.size[1]; addY++ {
			p := point{x: fabric.position[0] + addX, y: fabric.position[1] + addY}
			numOfOverlaps := takenFabric[p]
			if numOfOverlaps != 1 {
				return -1
			}
		}
	}
	return fabric.ID
}

// get the only one fabric that does not overlap with others (returning array in case
// there are more)
func getLoneFabric(fabricArr []elfFabric) []int {
	takenFabric := createFabricClaims(fabricArr)
	ids := []int{}
	// for each point check for all positions if they exist in onlyOneClaim
	for _, fabric := range fabricArr {
		id := checkPoints(&fabric, takenFabric)
		if id > 0 {
			ids = append(ids, id)
		}
	}
	return ids
}
