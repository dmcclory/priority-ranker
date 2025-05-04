package cmd

import (
	"slices"
	"math"
)

type WinRecord map[uint]map[uint]int64
type PScores map[uint]float64

func buildWinRecordFromVotes(votes []Vote, optionIds []uint) WinRecord {
	winRecord := make(WinRecord)
	for _, i := range optionIds {
		winRecord[i] = make(map[uint]int64)
		for _, j := range optionIds {
			if i != j {
				winRecord[i][j]= 0
			}
		}
	}

	for _, vote := range votes {
		winRecord[vote.WinnerId][vote.LoserId] += 1
	}
	return winRecord
}

func calcNumerator(wins WinRecord, pScores PScores, i uint, j uint) float64 {
	res := (float64(wins[i][j]) * pScores[j]) / (pScores[i] + pScores[j])
	if math.IsNaN(res) {
		return 0
	} else {
		return res
	}
}

func calcDenominator(wins WinRecord, pScores PScores, i uint, j uint) float64 {
	res := (float64(wins[j][i])) / (pScores[i] + pScores[j])
	if math.IsNaN(res) {
		return 0
	} else {
		return res
	}
}

func calcPScore(wins WinRecord, pScores PScores, i uint) float64 {
	var numeratorTotal float64
	var denominatorTotal float64

	for j, _ := range wins {
		numeratorTotal += calcNumerator(wins, pScores, i, j)
		denominatorTotal += calcDenominator(wins, pScores, i, j)
	}
	
	return numeratorTotal / denominatorTotal
}

func calcNewPScores(wins WinRecord, pScores PScores) PScores {
	newPScores := make(PScores)

	for k, score := range pScores {
		newPScores[k] = score
	}

	sortedKeys := []uint{}
	for k := range wins {
		sortedKeys = append(sortedKeys, k)
	}
	slices.Sort(sortedKeys)

	for _, i := range sortedKeys {
		newPScores[uint(i)] = calcPScore(wins, newPScores, uint(i))
	}

	return newPScores
}
