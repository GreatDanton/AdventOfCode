package main

import (
	"testing"
)

func Test_reactPolymer(t *testing.T) {
	testInputs := []struct {
		polymer         string
		expectedPolymer string
		units           int
	}{
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA", 10},
		{"daCaA", "daC", 3},
		{"CAacb", "b", 1},
		{"aaAAa", "a", 1},
		{"aA", "", 0},
		{"abBA", "", 0},
		{"abAB", "abAB", 4},
		{"aabAAB", "aabAAB", 6},
	}

	for _, test := range testInputs {
		polymer, units := alchemicalReduction(test.polymer)
		if polymer != test.expectedPolymer {
			t.Errorf("Initial: %v, Expected: %v, actual: %v", test.polymer, test.expectedPolymer, polymer)
		}
		if units != test.units {
			t.Errorf("Polymer units are not reported correctly. Expected: %v, actual: %v", test.units, units)
		}
	}
}

func Test_checkIfDestroyable(t *testing.T) {
	testInputs := []struct {
		a        rune
		b        rune
		expected bool
	}{
		{'a', 'b', false},
		{'A', 'B', false},
		{'A', 'b', false},
		{'a', 'a', false},
		{'A', 'A', false},
		{'a', 'A', true},
		{'A', 'a', true},
	}

	for _, test := range testInputs {
		destroyable := checkIfDestroyable(test.a, test.b)
		if destroyable != test.expected {
			t.Errorf("Runes %v and %v, destroyable: expected: %v, actual: %v", test.a, test.b, test.expected, destroyable)
		}
	}
}

func Test_removePolymer(t *testing.T) {
	testInputs := []struct {
		inputPolymer    string
		expectedPolymer string
		char            rune
	}{
		{"dabAcCaCBAcCcaDA", "dbcCCBcCcD", 'a'},
		{"dabAcCaCBAcCcaDA", "daAcCaCAcCcaDA", 'b'},
		{"dabAcCaCBAcCcaDA", "dabAaBAaDA", 'c'},
		{"dabAcCaCBAcCcaDA", "abAcCaCBAcCcaA", 'd'},
	}

	for _, test := range testInputs {
		removedPolymer := removeUnits(test.inputPolymer, test.char)
		if removedPolymer != test.expectedPolymer {
			t.Errorf("Expected: %v, actual: %v", test.expectedPolymer, removedPolymer)
		}
	}
}
