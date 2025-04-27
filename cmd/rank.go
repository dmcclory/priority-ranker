package cmd

type WinRecord map[int64]map[int64]int64
type PScores map[int64]float64

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
