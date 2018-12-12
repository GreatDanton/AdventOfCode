package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type star struct {
	x  int
	y  int
	vx int
	vy int
}

/* 10.day: https://adventofcode.com/2018/day/10 */
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// testing input (prints HI)
	/* 	lines = []string{
		"position=< 9,  1> velocity=< 0,  2>",
		"position=< 7,  0> velocity=<-1,  0>",
		"position=< 3, -2> velocity=<-1,  1>",
		"position=< 6, 10> velocity=<-2, -1>",
		"position=< 2, -4> velocity=< 2,  2>",
		"position=<-6, 10> velocity=< 2, -2>",
		"position=< 1,  8> velocity=< 1, -1>",
		"position=< 1,  7> velocity=< 1,  0>",
		"position=<-3, 11> velocity=< 1, -2>",
		"position=< 7,  6> velocity=<-1, -1>",
		"position=<-2,  3> velocity=< 1,  0>",
		"position=<-4,  3> velocity=< 2,  0>",
		"position=<10, -3> velocity=<-1,  1>",
		"position=< 5, 11> velocity=< 1, -2>",
		"position=< 4,  7> velocity=< 0, -1>",
		"position=< 8, -2> velocity=< 0,  1>",
		"position=<15,  0> velocity=<-2,  0>",
		"position=< 1,  6> velocity=< 1,  0>",
		"position=< 8,  9> velocity=< 0, -1>",
		"position=< 3,  3> velocity=<-1,  1>",
		"position=< 0,  5> velocity=< 0, -1>",
		"position=<-2,  2> velocity=< 2,  0>",
		"position=< 5, -2> velocity=< 1,  2>",
		"position=< 1,  4> velocity=< 2,  1>",
		"position=<-2,  7> velocity=< 2, -2>",
		"position=< 3,  6> velocity=<-1, -1>",
		"position=< 5,  0> velocity=< 1,  0>",
		"position=<-6,  0> velocity=< 2,  0>",
		"position=< 5,  9> velocity=< 1, -2>",
		"position=<14,  7> velocity=<-2,  0>",
		"position=<-3,  6> velocity=< 2, -1>",
	} */

	stars := parseInput(lines)
	displayStars(stars)
}

var maxWidth = 100
var maxHeight = 50
var maxIteration = 100000
var minMapSize = 100000000

// iteration where coordinates produced the smallest bounding box
var minMapWidthIteration = 0

func displayStars(stars []star) {
	seconds := 0
	for {
		seconds++
		currentWidth := width(stars)
		if currentWidth < minMapSize {
			minMapSize = currentWidth
			minMapWidthIteration = seconds
		}
		for i := 0; i < len(stars); i++ {
			star := stars[i]
			star.x = star.x + star.vx
			star.y = star.y + star.vy
			stars[i] = star
		}

		// minimum bounding box is achieved in 10244
		// seconds -1 will print the correct result
		if seconds == 10243 {
			printStars(stars)
		}

		if seconds > maxIteration {
			fmt.Println("iteration: ", minMapWidthIteration)
			break
		}
	}
}

func printStars(myStars []star) {
	stars := unique(myStars)
	// sort stars by y and sort by x ascending in both ways
	sort.Slice(stars, func(i, j int) bool {
		if stars[i].y < stars[j].y {
			return true
		}
		if stars[i].y > stars[j].y {
			return false
		}
		return stars[i].x < stars[j].x
	})
	fmt.Println(stars)

	iter := 0
	printed := []star{}
	for y := -maxHeight / 2; y < maxHeight/2; y++ {
		for x := -maxWidth / 2; x < maxWidth/2; x++ {
			if iter >= len(stars) {
				fmt.Printf(". ")
				continue
			}

			star := stars[iter]
			if star.y == y && star.x == x {
				printed = append(printed, star)
				fmt.Printf("x ")
				iter++
			} else { // such star does not exist
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("Len: ", len(stars), "Iter: ", iter, "Printed: ", len(printed))
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

func parseInput(input []string) []star {
	findNumbers := regexp.MustCompile(`-?\d+`)
	stars := []star{}
	for _, line := range input {
		parsedNumStrings := findNumbers.FindAllString(line, -1)
		if len(parsedNumStrings) < 4 {
			log.Fatal("Less than 4 parsed numbers")
		}
		n := convertToNum(parsedNumStrings)
		star := star{x: n[0], y: n[1], vx: n[2], vy: n[3]}
		stars = append(stars, star)
	}
	fmt.Println(stars[len(stars)-1])
	return stars
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(stars []star) (int, int) {
	maxX := stars[0].x
	maxY := stars[0].y

	for _, star := range stars {
		if star.x > maxX {
			maxX = star.x
		}
		if star.y > maxY {
			maxY = star.y
		}
	}
	return maxX, maxY
}

func min(stars []star) (int, int) {
	minX := stars[0].x
	minY := stars[0].y

	for _, star := range stars {
		if star.x < minX {
			minX = star.x
		}

		if star.y < minY {
			minY = star.y
		}
	}
	return minX, minY
}

func width(stars []star) int {
	minX, minY := min(stars)
	maxX, maxY := max(stars)
	return abs(minX-maxX) + abs(maxY-minY)
}

func unique(intSlice []star) []star {
	keys := make(map[star]bool)
	list := []star{}
	for _, entry := range intSlice {
		// reducing coordinates to move them more towards the center of the screen (instead of right down part)
		entry.x -= 170
		entry.y -= 170
		entry.vx = 0
		entry.vy = 0
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
