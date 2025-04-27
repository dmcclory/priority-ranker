package cmd

type WinRecord map[int64]map[int64]int64
type PScores map[int64]float64

func calcNumerator(wins WinRecord, pScores PScores, i int64, j int64) float64 {
	return (float64(wins[i][j]) * pScores[j]) / (pScores[i] + pScores[j])
}
