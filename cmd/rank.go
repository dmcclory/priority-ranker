package cmd

type WinRecord map[int64]map[int64]int64
type PScores map[int64]float64

func buildWinRecordFromVotes(votes []Vote, optionIds []int64) WinRecord {
	winRecord := make(WinRecord)
	for _, i := range optionIds {
		winRecord[i] = make(map[int64]int64)
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

func calcNumerator(wins WinRecord, pScores PScores, i int64, j int64) float64 {
	return (float64(wins[i][j]) * pScores[j]) / (pScores[i] + pScores[j])
}

func calcDenominator(wins WinRecord, pScores PScores, i int64, j int64) float64 {
	return (float64(wins[j][i])) / (pScores[i] + pScores[j])
}

func calcPScore(wins WinRecord, pScores PScores, i int64) float64 {
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

	// fmt.Println(newPScores)

	for i, _ := range wins {
		newPScores[i] = calcPScore(wins, newPScores, i)
	}
	
	return newPScores
}
