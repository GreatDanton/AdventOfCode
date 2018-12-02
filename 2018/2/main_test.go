package main

import "testing"

func Test_repeatedLetters(t *testing.T) {
	testInputs := []struct {
		input     string
		twoChar   bool
		threeChar bool
	}{
		{"abcdef", false, false},
		{"bababc", true, true},
	}

	for _, test := range testInputs {
		actualTwoChar, actualThreeChar := repeatedLetters(test.input)
		if actualTwoChar != test.twoChar || actualThreeChar != test.threeChar {
			t.Errorf("calculateWordChecksum(%v) expected: twoChar: %v threeChar: %v, actual: twoChar: %v threeChar: %v",
				test.input, test.twoChar, test.threeChar, actualTwoChar, actualThreeChar)
		}
	}
}

func Test_calculateChecksum(t *testing.T) {
	testInputs := []struct {
		input    []string
		expected int
	}{
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}

	for _, test := range testInputs {
		checksum := calculateChecksum(test.input)
		if checksum != test.expected {
			t.Errorf("Checksum expected: %v, actual: %v", test.expected, checksum)
		}
	}
}

func Test_oneCharDifference(t *testing.T) {
	testInputs := []struct {
		firstWord  string
		secondWord string
		expected   bool
	}{
		{"abc", "abc", true},
		{"abc", "abd", true},
		{"abc", "bbc", true},
		{"abc", "abcd", false},
	}

	for _, test := range testInputs {
		isOneCharDiff := oneCharDifference(test.firstWord, test.secondWord)
		if isOneCharDiff != test.expected {
			t.Errorf("Words %s %s are not one character apart, expected: %v actual: %v", test.firstWord, test.secondWord, test.expected, isOneCharDiff)
		}
	}
}

func Test_getSameIndexCharacters(t *testing.T) {
	testInputs := []struct {
		firstWord  string
		secondWord string
		expected   string
	}{
		{"abcd", "abcD", "abc"},
		{"Abc", "abc", "bc"},
		{"abc", "abcd", "abc"},
		{"abcd", "abc", "abc"},
		{"abcd", "efcb", "c"},
	}

	for _, test := range testInputs {
		sameChars := getSameIndexCharacters(test.firstWord, test.secondWord)

		if sameChars != test.expected {
			t.Errorf("getSameCharacters(%s, %s) expected: %s, actual: %s", test.firstWord, test.secondWord, test.expected, sameChars)
		}
	}
}
