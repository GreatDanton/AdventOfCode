package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guardNote struct {
	t    string
	note string
}

// 4.day: https://adventofcode.com/2018/day/4
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parsedLines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLines = append(parsedLines, line)
	}

	guardNotes := parseInput(parsedLines)
	guardTimes := parseGuardNotes(guardNotes)
	sleepyGuard, maxSleep, maxMinute := getSleepyGuard(guardTimes)
	fmt.Printf("Sleepy guard id: %d, max sleep: %d, minute asleep: %d\n", sleepyGuard, maxSleep, maxMinute)
	fmt.Println("Sleepy guard: ", sleepyGuard*maxMinute)

	fmt.Println()
	fmt.Println("=== Part 2 ===")
	guardID, minute := mostFrequentlyAsleepGuard(guardTimes)
	fmt.Printf("Guard id: %d, minute: %d\n", guardID, minute)
	fmt.Println("Most frequently asleep guard:", guardID*minute)
}

// part 1
//--------------------------------------------
func parseInput(input []string) []guardNote {
	notes := []guardNote{}
	regex := regexp.MustCompile(`\[(.*)\](.*)`)
	for _, line := range input {
		outArr := regex.FindAllStringSubmatch(line, -1)
		for i := range outArr {
			// we can already sort by string without parsing it into time format
			dateStr := outArr[i][1]
			action := outArr[i][2]
			notes = append(notes, guardNote{t: dateStr, note: action})
		}
	}
	// order notes by time ascending
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].t < notes[j].t
	})
	return notes
}

func parseGuardNotes(guardNotes []guardNote) map[int][]int {
	const maxMinutes int = 60
	guards := map[int][]int{}
	currentGuard := -1
	wakesUpMin := -1
	fallsAsleepMin := -1
	getNum := regexp.MustCompile(`\d+`)

	for _, n := range guardNotes {
		if currentGuard == -1 && !strings.Contains(n.note, "Guard") {
			// skip starting lines where we don't know which guard was tracked
			continue
		}

		if strings.Contains(n.note, "Guard") {
			guardNum, err := strconv.Atoi(getNum.FindString(n.note))
			if err != nil {
				fmt.Println(err)
				continue
			}
			_, exists := guards[guardNum]
			if !exists {
				// add copy of a null arr
				guards[guardNum] = make([]int, maxMinutes)
			}
			currentGuard = guardNum
			wakesUpMin = -1
			fallsAsleepMin = -1
		} else if strings.Contains(n.note, "wakes") { // guard wakes up
			wakesUpMin = parseMinute(n.t)
			updateSleepingTime(guards, currentGuard, fallsAsleepMin, wakesUpMin)
		} else if strings.Contains(n.note, "falls") { // guard falls asleep
			fallsAsleepMin = parseMinute(n.t)
		}
	}
	return guards
}

func parseMinute(time string) int {
	t, err := strconv.Atoi(strings.Split(time, ":")[1])
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return t
}

func updateSleepingTime(guards map[int][]int, guard int, fallsAsleepMin int, wakesUpMin int) {
	// early return on invalid values
	if fallsAsleepMin < 0 || wakesUpMin < 0 {
		return
	}
	v := guards[guard]
	sleep := wakesUpMin - fallsAsleepMin
	for start := fallsAsleepMin; start < fallsAsleepMin+sleep; start++ {
		v[start]++
	}
}

func getSleepyGuard(guardTimes map[int][]int) (int, int, int) {
	maxSleep := -1
	sleepyMinute := -1
	sleepyGuard := -1
	for key, sleepingTime := range guardTimes {
		sum, minute := sumAndMaxValIndex(sleepingTime)
		if sum > maxSleep {
			maxSleep = sum
			sleepyGuard = key
			sleepyMinute = minute
		}
	}

	return sleepyGuard, maxSleep, sleepyMinute
}

// part 2:
// which guard is most frequently asleep on the same minute
// return guard id, most frequently asleep minute
func mostFrequentlyAsleepGuard(guardTimes map[int][]int) (int, int) {
	maxGuard := -1
	minute := -1
	maxSleep := -1
	for guard, sleepTimes := range guardTimes {
		sleep, min := maxValAndIndex(sleepTimes)
		if sleep > maxSleep {
			maxSleep = sleep
			minute = min
			maxGuard = guard
		}
	}

	return maxGuard, minute
}

// helper functions (reusable functions)
// ------------------------------------------
// get sum of the array and index of the maximum value in the array
func sumAndMaxValIndex(arr []int) (int, int) {
	sum := 0    // minutes of sleep
	minute := 0 // most asleep minute
	maxValIndex := 0

	for i := 0; i < len(arr); i++ {
		element := arr[i]
		sum += element
		if element > maxValIndex {
			maxValIndex = element
			minute = i
		}
	}
	return sum, minute
}

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

// parsing string into Time example
func parsingTime(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02 15:04", dateStr)
	if err != nil {
		fmt.Println(err)
	}
	return t
}
