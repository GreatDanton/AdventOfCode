package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 2. day: https://adventofcode.com/2018/day/2
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

	checksum := calculateChecksum(input)
	fmt.Printf("Final checksum is %d\n", checksum)

	filteredBoxes := wordsWithOneCharDifference(input)
	fmt.Println("Boxes with similar ids:", filteredBoxes)
	fmt.Println("")

	boxPair := filteredBoxes[0]
	sameChars := getSameIndexCharacters(boxPair[0], boxPair[1])

	fmt.Println("Same characters are: ")
	fmt.Println(sameChars)
}

// Scans the file and appends each line into an array of lines
func readInput(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)
	input := []string{}
	for scanner.Scan() {
		boxID := scanner.Text()
		input = append(input, boxID)
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return input, nil
}

// 1. part
// ----------------------------------------
// calculate checksum of the box ids
func calculateChecksum(boxIds []string) int {
	amountOfTwos := 0
	amountOfThrees := 0
	for _, line := range boxIds {
		twoChar, threeChar := repeatedLetters(line)
		if twoChar {
			amountOfTwos++
		}
		if threeChar {
			amountOfThrees++
		}
	}

	return amountOfTwos * amountOfThrees
}

// Find letters that are repeated in the string
//
// first return value is true if 2 same characters exist in the word
// second return value is true if 3 same characters exist in the word
func repeatedLetters(boxID string) (bool, bool) {
	letters := map[rune]int{}

	for _, letter := range boxID {
		letters[letter]++
	}

	twoChar := false
	threeChar := false
	for _, value := range letters {
		switch value {
		case 2:
			twoChar = true
			break
		case 3:
			threeChar = true
			break
		default:
			// do nothing
		}
	}

	return twoChar, threeChar
}

// 2. part:
// ----------------------------------------
// at most one character difference between the strings
func oneCharDifference(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}

	numOfDifferentChars := 0
	for i := range first {
		firstChar := first[i]
		secondChar := second[i]

		if firstChar != secondChar {
			numOfDifferentChars++
		}

		if numOfDifferentChars > 1 {
			return false
		}
	}

	return true
}

// tests for at most one char difference between every word in the input fields
func wordsWithOneCharDifference(input []string) [][]string {
	filtered := [][]string{}
	for i, firstWord := range input {
		for j, secondWord := range input {
			if i == j {
				continue
			}

			// this will actually produce duplicated values which we don't want
			// TODO: get rid of duplicates
			if oneCharDifference(firstWord, secondWord) {
				pair := make([]string, 2, 2)
				pair[0] = firstWord
				pair[1] = secondWord
				filtered = append(filtered, pair)
			}
		}
	}

	return filtered
}

// get the characters that are present in both strings on the same index
func getSameIndexCharacters(first string, second string) string {
	if len(first) > len(second) {
		first, second = second, first
	}

	sameChars := []byte{}
	for i := range first {
		firstChar := first[i]
		secondChar := second[i]
		if firstChar == secondChar {
			sameChars = append(sameChars, firstChar)
		}
	}

	return string(sameChars)
}
