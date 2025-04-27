package cmd

import (
  "testing"
	"math"
)

func nearlyEqual(a float64, b float64) bool {
	return math.Abs(a - b) <= 0.00001
}

func buildExample() WinRecord {

	return WinRecord{
		1: {2: 2, 3: 0, 4: 1},
		2: {1: 3, 3: 5, 4: 0},
		3: {1: 0, 2: 3, 4: 1},
		4: {1: 4, 2: 0, 3: 3},
	}
}

func buildInitialPScores() PScores {

	return PScores{
		1: 1,
		2: 1,
		3: 1,
		4: 1,
	}
}

// I was going to make a type that would hold the pairs i & j
// and put that on an "input" prop, but realized it would be easier to just
// do i & j as separate fields on the test case itself
// type IndexPair = [2]int64

func TestCalculateNumeratorTable(t *testing.T) {
	// based on: https://go.dev/wiki/TableDrivenTests#using-a-map-to-store-test-cases
	tests := map[string]struct {
		i int64
		j int64
		result float64
	} {
		"numerator with 1, 2": {
			i: 1, j: 2, result: 1,
		},
		"numerator with 1, 3": {
			i: 1, j: 3, result: 0,
		},
		"numerator with 1, 4": {
			i: 1, j: 4, result: 0.5,
		},
	}

	wins := buildExample()
	pScores := buildInitialPScores()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got, expected := calcNumerator(wins, pScores, test.i, test.j), test.result; got != expected {
				t.Fatalf("calcNumerator with %d and %d returned %f, expected %f", test.i, test.j, got, expected)
			}
		})
	}
}

func TestCalculateDenominatorTable(t *testing.T) {
	tests := map[string]struct {
		i int64
		j int64
		result float64
	} {
		"denominator with 1, 2": {
			i: 1, j: 2, result: 1.5,
		},
		"denominator with 1, 3": {
			i: 1, j: 3, result: 0,
		},
		"denominator with 1, 4": {
			i: 1, j: 4, result: 2,
		},
	}

	wins := buildExample()
	pScores := buildInitialPScores()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got, expected := calcDenominator(wins, pScores, test.i, test.j), test.result; got != expected {
				t.Fatalf("calcDenominator with %d and %d returned %f, expected %f", test.i, test.j, got, expected)
			}
		})
	}
}

func TestCalculateIndividualPScore(t *testing.T) {
	tests := map[string]struct {
		i int64
		pScores PScores
		result float64
	} {
		"pScore for cell 1": {
			i: 1, pScores: PScores{1: 1, 2: 1, 3: 1, 4: 1}, result: 0.428571,
		},
		"pScore for cell 2": {
			i: 2, pScores: PScores{1: 0.428571, 2: 1, 3: 1, 4: 1}, result: 1.172413,
		},
	}

	wins := buildExample()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got, expected := calcPScore(wins, test.pScores, test.i), test.result; !nearlyEqual(got, expected) {
				t.Fatalf("calcPScore with %d failed, expected %f, got %f", test.i, test.result, got)
			}
		})
	}
}

func TestCalculateNewPScoresTable(t *testing.T) {
	tests := map[string]struct {
		i int64
		result float64
	} {
		"new pScore for cell 1": { i: 1, result: 0.428571, },
		"new pScore for cell 2": { i: 2, result: 1.172413, },
		"new pScore for cell 3": { i: 3, result: 0.557411, },
		"new pScore for cell 4": { i: 4, result: 1.694167, },
	}

	wins := buildExample()
	pScores := buildInitialPScores()

	t.Parallel()
	for name, test:= range tests {
		t.Run(name, func(t *testing.T) {
			if got, expected := calcNewPScores(wins, pScores), test.result; !nearlyEqual(got[test.i], expected) {
				t.Fatalf("calcNewPScores with %d failed, expected %f, got %f", test.i, test.result, got[test.i])
			}
		})
	}
}

func TestCalculateNewPScores(t *testing.T) {
	wins := buildExample()
	pScores := buildInitialPScores()

	newPScores := calcNewPScores(wins, pScores)

	if !nearlyEqual(newPScores[1], 0.428571) {
		t.Errorf("Expected %f, got %f\n", 0.428571, newPScores[1])
	}

	if !nearlyEqual(newPScores[2], 1.172414) {
		t.Errorf("Expected %f, got %f\n", 1.172414, newPScores[2])
	}

	if !nearlyEqual(newPScores[3], 0.557411) {
		t.Errorf("Expected %f, got %f\n", 0.557411, newPScores[3])
	}

	if !nearlyEqual(newPScores[4], 1.694167) {
		t.Errorf("Expected %f, got %f\n", 1.694167, newPScores[4])
	}
}
