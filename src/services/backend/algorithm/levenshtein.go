package algorithm

import (
	"fmt"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func searchQuestion(pattern string, questions []string) (int, float64) {
	// Cari pertanyaan yang cocok dengan pattern menggunakan KMP
	for i, question := range questions {
		if KnuthMorrisPratt(question, pattern) != -1 {
			return i, 1.0
		}
	}

	// Cari pertanyaan yang mirip dengan pattern menggunakan Levenshtein
	minDist := len(pattern)
	var index int = -1
	for i, question := range questions {
		dist := levenshtein.DistanceForStrings([]rune(pattern), []rune(question), levenshtein.DefaultOptions)
		similarity := 1.0 - (float64(dist) / float64(len(question)))
		if similarity == 1.0 {
			return i, 1.0
		} else if similarity > 0.7 && similarity > float64(minDist)/float64(len(question)) {
			minDist = dist
			index = i
		}
	}

	// Jika jarak terlalu jauh, anggap tidak ada pertanyaan yang cocok
	if minDist > len(pattern)/2 {
		index = -1
		fmt.Println("Tidak ada pertanyaan yang cocok")
	}

	// Hitung persentase kemiripan
	var similarity float64 = 0.0
	if index != -1 {
		similarity = 1.0 - (float64(minDist)/float64(len(questions[index])))*100.0
		if similarity < 50.0 {
			index = -1
			fmt.Println("Tidak ada pertanyaan yang cocok")
		}
	}

	return index, similarity
}
