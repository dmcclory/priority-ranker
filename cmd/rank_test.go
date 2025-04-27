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

func TestCalculateNumerator(t *testing.T) {
	wins := buildExample()
	pScores := buildInitialPScores()

	cell1 := calcNumerator(wins, pScores, 1, 2)

	if cell1 != 1 {
		t.Errorf("was expecting something different man!")
	}

	cell2 := calcNumerator(wins, pScores, 1, 3)

	if cell2 != 0 {
		t.Errorf("was expecting something different man!")
	}

	cell3 := calcNumerator(wins, pScores, 1, 4)

	if cell3 != 0.5 {
		t.Errorf("was expecting something different man!")
	}
}

func TestCalculateDenominator(t *testing.T) {
	wins := buildExample()
	pScores := buildInitialPScores()

	cell1 := calcDenominator(wins, pScores, 1, 2)

	if cell1 != 1.5 {
		t.Errorf("was expecting something different man!")
	}

	cell2 := calcDenominator(wins, pScores, 1, 3)

	if cell2 != 0 {
		t.Errorf("was expecting something different man!")
	}

	cell3 := calcDenominator(wins, pScores, 1, 4)

	if cell3 != 2 {
		t.Errorf("was expecting something different man!")
	}
}

func TestCalculatePScore(t *testing.T) {
	wins := buildExample()
	pScores := buildInitialPScores()

	cell1 := calcPScore(wins, pScores, 1)

	if !nearlyEqual(cell1, 0.428571) {
		t.Errorf("was expecting something different man!, got: %f", cell1)
	}
}

func TestCalculatePScoreWithSomeOtherEntries(t *testing.T) {
	wins := buildExample()
	pScores := buildInitialPScores()

	pScores[1] = 0.428571

	cell2 := calcPScore(wins, pScores, 2)

	if !nearlyEqual(cell2, 1.172413) {
		t.Errorf("was expecting something different man!, got: %f", cell2)
	}
}
