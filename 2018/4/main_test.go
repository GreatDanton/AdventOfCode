package main

import "testing"

func Test_getMaxValAndIndex(t *testing.T) {
	testInputs := []struct {
		input          []int
		expectedSum    int
		expectedMinute int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 21, 5},
		{[]int{0, 5, 1, 0}, 6, 1},
	}

	for _, test := range testInputs {
		sum, minute := sumAndMaxValIndex(test.input)
		if sum != test.expectedSum {
			t.Errorf("Sum: Expected: %d, Actual: %d", test.expectedSum, sum)
		}
		if minute != test.expectedMinute {
			t.Errorf("Minute: Expected: %d, Actual: %d", test.expectedMinute, minute)
		}
	}
}

func Test_parseInput(t *testing.T) {
	testInputs := []struct {
		input    []string
		expected int
	}{
		{[]string{
			"[1518-11-01 00:00] Guard #10 begins shift",
			"[1518-11-01 00:05] falls asleep",
			"[1518-11-01 00:25] wakes up",
			"[1518-11-01 00:30] falls asleep",
			"[1518-11-01 00:55] wakes up",
			"[1518-11-01 23:58] Guard #99 begins shift",
			"[1518-11-02 00:40] falls asleep",
			"[1518-11-02 00:50] wakes up",
			"[1518-11-03 00:05] Guard #10 begins shift",
			"[1518-11-03 00:24] falls asleep",
			"[1518-11-03 00:29] wakes up",
			"[1518-11-04 00:02] Guard #99 begins shift",
			"[1518-11-04 00:36] falls asleep",
			"[1518-11-04 00:46] wakes up",
			"[1518-11-05 00:03] Guard #99 begins shift",
			"[1518-11-05 00:45] falls asleep",
			"[1518-11-05 00:55] wakes up",
		}, 240},
	}

	for _, test := range testInputs {
		guardNotes := parseInput(test.input)
		guardTimes := parseGuardNotes(guardNotes)
		sleepyGuard, _, maxMinute := getSleepyGuard(guardTimes)
		guard := sleepyGuard * maxMinute
		if guard != test.expected {
			t.Errorf("Wrong type of guard, Expected: %d, actual: %d", test.expected, guard)
		}
	}
}
