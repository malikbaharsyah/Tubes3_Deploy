package algorithm

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func searchQuestion(pattern string, questions []string, param int) ([]string, []int) {
	if param == 0 {
		for i, question := range questions {
			if KnuthMorrisPratt(question, pattern) != -1 {
				return []string{question}, []int{i}
			}
		}
		return lv(pattern, questions)
	} else {
		for i, question := range questions {
			if BoyerMooreAlgorithm(question, pattern) != -1 {
				return []string{question}, []int{i}
			}
		}
		return lv(pattern, questions)
	}
}

func lv(pattern string, questions []string) ([]string, []int) {
	type matrix struct {
		index int
		sim   float64
	}
	matrixSim := make([]matrix, len(questions))
	for i, question := range questions {
		matrixSim[i].index = i
		dist := levenshtein.DistanceForStrings([]rune(pattern), []rune(question), levenshtein.DefaultOptions)
		matrixSim[i].sim = float64(dist) / float64(len(pattern))
	}
	for i := 0; i < len(matrixSim); i++ {
		for j := i + 1; j < len(matrixSim); j++ {
			if matrixSim[i].sim > matrixSim[j].sim {
				temp := matrixSim[i]
				matrixSim[i] = matrixSim[j]
				matrixSim[j] = temp
			}
		}
	}
	highestSim := matrixSim[0].sim
	if highestSim > 0.8 {
		return []string{questions[matrixSim[0].index]}, []int{matrixSim[0].index}
	}
	var top3 []float64
	var top3Index []int
	for i := 0; i < 3; i++ {
		top3 = append(top3, matrixSim[i].sim)
		top3Index = append(top3Index, matrixSim[i].index)
	}
	for i := 0; i < len(top3); i++ {
		if top3[i] > 0.4 && top3[i] < 0.8 {
			return []string{questions[top3Index[0]], questions[top3Index[1]], questions[top3Index[2]]}, []int{top3Index[0], top3Index[1], top3Index[2]}
		}
	}
	return []string{"Tidak ada pertanyaan yang cocok"}, []int{-1}
}
