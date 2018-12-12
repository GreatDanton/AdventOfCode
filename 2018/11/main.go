package main

import (
	"fmt"
	"log"
	"strconv"
)

// 11.day: https://adventofcode.com/2018/day/11
func main() {
	const gridID = 6303
	fmt.Println(coordinatePowerLever(3, 5, 8))

	grid := createGrid(gridID, 300)
	powerLevel, x, y := getMaxPowerLevelSquare(grid, 3)
	fmt.Printf("Power level: %d, x: %d, y: %d\n", powerLevel, x, y)

	maxLevel := 0
	maxX := 0
	maxY := 0
	maxSquareSize := 0
	// this is brute force approach it could be done better
	for squareSize := 1; squareSize < 301; squareSize++ {
		level, x, y := getMaxPowerLevelSquare(grid, squareSize)
		if level > maxLevel {
			maxLevel = level
			maxX = x
			maxY = y
			maxSquareSize = squareSize
			fmt.Println(x, y, maxSquareSize)
		}
	}
	fmt.Printf("Max level of all squares: x,y,squareSize %d,%d,%d\n", maxX, maxY, maxSquareSize)
}

func createGrid(gridID int, gridSize int) [][]int {
	grid := [][]int{}
	for y := 0; y < gridSize; y++ {
		xRow := make([]int, gridSize)
		for x := 0; x < gridSize; x++ {
			xRow[x] = coordinatePowerLever(x+1, y+1, gridID)
		}
		grid = append(grid, xRow)
	}
	return grid
}

func getMaxPowerLevelSquare(grid [][]int, squareSize int) (int, int, int) {
	maxX := 1
	maxY := 1
	maxSquareLevel := squarePowerLevel(1, 1, squareSize, grid)
	maxSize := len(grid) - squareSize + 1
	for y := 0; y < maxSize; y++ {
		for x := 0; x < maxSize; x++ {
			squareLevel := squarePowerLevel(x, y, squareSize, grid)
			if squareLevel > maxSquareLevel {
				maxSquareLevel = squareLevel
				maxX = x + 1
				maxY = y + 1
			}
		}
	}
	return maxSquareLevel, maxX, maxY
}

// 3x3 square power level
func squarePowerLevel(x, y, squareSize int, grid [][]int) int {
	squarePowerLevel := 0
	for dy := 0; dy < squareSize; dy++ {
		for dx := 0; dx < squareSize; dx++ {
			newY := y + dy
			newX := x + dx
			if newX > len(grid)-1 || newY > len(grid)-1 {
				return squarePowerLevel
			}
			squarePowerLevel += grid[newY][newX]
		}
	}
	return squarePowerLevel
}

func coordinatePowerLever(x, y, gridID int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += gridID
	powerLevel *= rackID
	powerLevel = getHundredsOrZero(powerLevel)
	powerLevel -= 5
	return powerLevel
}

func getHundredsOrZero(num int) int {
	strNum := strconv.Itoa(num)
	if len(strNum) >= 3 {
		position := len(strNum) - 1 - 2
		n, err := strconv.Atoi(string(strNum[position]))
		if err != nil {
			log.Fatal(err)
		}
		return n
	}
	return 0
}
