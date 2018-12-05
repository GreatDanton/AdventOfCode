package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Trim(string(bytes), "\n")
	_, units := alchemicalReduction(input)
	fmt.Println("Units of polymer", units)

	// part 2
	shortestUnits, char := improvedPolymer(input)
	fmt.Printf("Shortest units by removing: %s units: %d\n", string(char), shortestUnits)
}

func alchemicalReduction(initialPolymer string) (string, int) {
	destroyed := 0
	polymer := initialPolymer
	for {
		/* 		fmt.Println(polymer) */
		polymer, destroyed = reactPolymer(polymer)
		if destroyed == 0 {
			break
		}
	}
	return polymer, len(polymer)
}

// destroy characters of different polarity (i.e: Aa - gets destroyed, aa - stays, Ab - stays)
func reactPolymer(polymer string) (string, int) {
	destroyed := 0
	endPolymer := []rune{}
	for i := 0; i < len(polymer); i++ {
		if i < len(polymer)-1 {
			firstChar := rune(polymer[i])
			secondChar := rune(polymer[i+1])

			destroyable := checkIfDestroyable(firstChar, secondChar)
			if destroyable {
				destroyed++
				i++
			} else {
				endPolymer = append(endPolymer, firstChar)
			}
		} else {
			// we are on the last character, can't compare it with any other char
			endPolymer = append(endPolymer, rune(polymer[i]))
		}
	}

	return string(endPolymer), destroyed
}

func checkIfDestroyable(a, b rune) bool {
	if !strings.EqualFold(string(a), string(b)) {
		return false
	}

	// check if one of them is uppercase and the other not
	if (unicode.IsLower(a) && !unicode.IsLower(b)) || !unicode.IsLower(a) && unicode.IsLower(b) {
		return true
	}
	// two upper case chars
	return false
}

// part 2
// ------------------------------------
func removeUnits(polymer string, char rune) string {
	finder := fmt.Sprintf("%s|%s", string(char), string(unicode.ToUpper(char)))
	regex := regexp.MustCompile(finder)
	improvedPolymer := regex.ReplaceAllString(polymer, "")
	return improvedPolymer
}

func improvedPolymer(polymer string) (int, rune) {
	minLength := len(polymer)
	minChar := ' '
	// remove all letters
	allLetters := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range allLetters {
		improvedPolymer := removeUnits(polymer, char)
		_, units := alchemicalReduction(improvedPolymer)
		fmt.Println(units, string(char))
		if units < minLength {
			minLength = units
			minChar = char
		}
	}
	return minLength, minChar
}
