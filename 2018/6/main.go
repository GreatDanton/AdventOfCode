package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x int
	y int
}

type boundaries struct {
	minX int
	maxX int
	minY int
	maxY int
}

// day 6: https://adventofcode.com/2018/day/6
// my eyes are bleeding, but it produces the correct result ;)
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		// TODO: handle error
	}
	points := turnInputToPoints(lines)
	// part one
	// set map -> min x, max x, min y, max y | that defines how large is our map
	xPoints := make([]int, len(points))
	yPoints := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		p := points[i]
		xPoints[i] = p.x
		yPoints[i] = p.y
	}
	// map boundaries
	minX := min(xPoints)
	maxX := max(xPoints)
	minY := min(yPoints)
	maxY := max(yPoints)
	dx := maxX - minX
	dy := maxY - minY
	mapBoundaries := boundaries{minX: minX, maxX: maxX, minY: minY, maxY: maxY}

	fmt.Println(mapBoundaries)
	fmt.Println(dx, dy)
	//
	maxR := max([]int{dx, dy})
	// this land does not exist, out of map point so it doesn't collide with any map?
	noMansLand := point{x: minX - 1, y: minY - 1}
	fmt.Println("No mans land", noMansLand)
	infiniteSources := map[point]int{noMansLand: 0}
	// land : source coordinates
	landClaims := map[point]point{}
	finishedCalculations := map[point]int{}

	for r := 1; r < maxR; r++ {
		// land coordinate: array of source coordinates
		thisRoundClaims := map[point][]point{}
		for _, source := range points {
			_, finished := finishedCalculations[source]
			if finished {
				continue
			}

			virtualPoints := claimVirtualPoints(source, r)
			// If virtual points falls outside of map, then source is one of the infinite sources
			existCounter := 0
			for _, land := range virtualPoints {
				// check if land is not already claimed, if it's not append that source
				// to the sources that claim this land in this round
				_, exists := landClaims[land]
				if exists {
					existCounter++
					continue
				}

				outside := outsideOfMap(mapBoundaries, land.x, land.y)
				if outside {
					infiniteSources[source]++
					continue
				}

				// append source to the sources that claim this land
				thisRoundClaims[land] = append(thisRoundClaims[land], source)
			}

			// all points around this source were already taken, stop iterating
			// this point
			if existCounter >= len(virtualPoints) {
				finishedCalculations[source]++
			}
		}

		for land, sources := range thisRoundClaims {
			// if there is only one source to claim the land, assign that source
			// to the land
			if len(sources) == 1 {
				source := sources[0]
				landClaims[land] = source
			} else if len(sources) > 1 {
				// if multiple sources are claiming the land, make it a no mans land
				// landClaims[noMansLand] = append(landClaims[noMansLand], land)
				landClaims[land] = noMansLand
			}
		}
	}

	// looping finished check the results
	// source: amount of land
	sourcesClaims := map[point]int{}
	for _, source := range landClaims {
		sourcesClaims[source]++
	}

	// loop through keys of claims and pick only the one that are not in infinite sources
	maxSource := noMansLand
	maxClaims := 0
	for source, claims := range sourcesClaims {
		_, exists := infiniteSources[source]
		if !exists && maxClaims < claims {
			maxSource = source
			maxClaims = claims
		}
	}

	fmt.Println("max source: ", maxSource)
	fmt.Println("max claims: ", maxClaims)
	fmt.Println(sourcesClaims)

	// part 2
	// distance less
	viablePoints := []point{}
	maxDistance := 10000
	for land := range landClaims {
		distances := 0
		for _, s := range points {
			d := distance(land, s)
			distances += d
		}

		if distances < maxDistance {
			viablePoints = append(viablePoints, land)
		}
	}

	fmt.Println("Viable points", len(viablePoints))
}

func turnInputToPoints(input []string) []point {
	points := []point{}
	for _, line := range input {
		x := 0
		y := 0
		out, err := fmt.Sscanf(line, "%d, %d", &x, &y)
		if err != nil {
			log.Fatal(err)
		}

		if out != 2 {
			log.Fatal("Wrong number of parsed arguments", out)
		}

		points = append(points, point{x: x, y: y})
	}

	return points
}

// checks if point is outside of map. True if one of its coordinates is outside of
// map and false if it's inside of map
func outsideOfMap(m boundaries, x, y int) bool {
	if x < m.minX || x > m.maxX {
		return true
	}

	if y < m.minY || y > m.maxY {
		return true
	}

	return false
}

// part 2
func distance(land point, source point) int {
	return abs(land.x-source.x) + abs(land.y-source.y)
}

// helper functions
// helper function returning maximum value from the array and it's respective index
func maxValAndIndex(arr []int) (int, int) {
	iterator := 0
	maxVal := arr[iterator]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxVal {
			iterator = i
			maxVal = arr[iterator]
		}
	}
	return maxVal, iterator
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// min max functions
func min(arr []int) int {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func max(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func contains(points []point, checkPoint point) bool {
	for _, p := range points {
		if p == checkPoint {
			return true
		}
	}
	return false
}

func claimVirtualPoints(source point, radius int) []point {
	left := point{x: source.x - radius, y: source.y}
	right := point{x: source.x + radius, y: source.y}
	top := point{x: source.x, y: source.y + radius}
	bot := point{x: source.x, y: source.y - radius}

	leftTop := createSides(left, top)
	rightTop := createSides(top, right)
	leftBot := createSides(left, bot)
	rightBot := createSides(bot, right)

	allPoints := []point{left, right, top, bot}
	allPoints = append(allPoints, leftTop...)
	allPoints = append(allPoints, rightTop...)
	allPoints = append(allPoints, leftBot...)
	allPoints = append(allPoints, rightBot...)
	return allPoints
}

func createSides(p1 point, p2 point) []point {
	points := []point{}
	k := (p2.y - p1.y) / (p2.x - p1.x)
	// y = k*x + n
	n := p2.y - k*p2.x
	// remove starting point and ending point
	for x := p1.x + 1; x < p2.x; x++ {
		y := k*x + n
		newPoint := point{x: x, y: y}
		points = append(points, newPoint)
	}
	return points
}
