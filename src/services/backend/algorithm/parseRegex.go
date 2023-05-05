package algorithm

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseInput(input string, listOfQuestion []string, param int) []string {
	patternDate := "^(?i)(Hari\\s\\d{1,2}/\\d{1,2}/\\d{4}|\\d{1,2}/\\d{1,2}/\\d{4})\\??$"
	regDate := regexp.MustCompile(patternDate)

	patternMath := `^(Hasil dari\s*|Hasil\s*)?-?\s*(\(\s*)*\s*-?\s*\d+(\.\d+)?(\s*\))*\s*(\s*[-+*/]\s*(\(\s*)*\s*-?\s*(\(\s*)*\d+(\.\d+)?(\s*\))*\s*)*\s*(\?\s*|\s\?\s*)?$`
	regCal := regexp.MustCompile(patternMath)

	patternAdd := `(?i)^tambahkan pertanyaan\s(.+)\sdengan jawaban\s(.+)$`
	regAdd := regexp.MustCompile(patternAdd)

	patternDel := `(?i)^hapus pertanyaan\s(.+)$`
	regDel := regexp.MustCompile(patternDel)

	if regDate.MatchString(input) {
		fmt.Println("Date")
		getInput := getDate(input)
		fmt.Println(getInput)
		date := Calendar(getInput)
		return []string{"kalender", "Hari " + date}
	} else if regCal.MatchString(input) {
		fmt.Println("Calculator")
		getInput := getCalculator(input)
		math := Calculator(getInput)
		return []string{"kalkulator", "Hasilnya adalah " + math}
	} else if regAdd.MatchString(input) {
		fmt.Println("Add Question")
		question, answer := extractQuestionAnswer(input)
		return []string{"tambah", question, answer}
	} else if regDel.MatchString(input) {
		fmt.Println("Delete Question")
		question := getQuestionDeleteCommand(input)
		return []string{"hapus", question}
	} else {
		fmt.Println("Question")
		question, index := searchQuestion(input, listOfQuestion, param)
		rekomendasi := []string{"rekomendasi"}
		if index[0] != -1 {
			rekomendasi = append(rekomendasi, "Apakah maksud Anda: ")
			for i := 0; i < len(question); i++ {
				rekomendasi = append(rekomendasi, question[i])
			}
		} else {
			rekomendasi = append(rekomendasi, "Saya tidak dapat menjawab pertanyaan Anda.")
		}
		return rekomendasi
	}
}

func getDate(input string) string {
	patt := "^(?i)(Hari\\s)?(\\d{1,2}/\\d{1,2}/\\d{4})\\??$"
	reg := regexp.MustCompile(patt)
	matches := reg.FindStringSubmatch(input)
	if len(matches) == 3 {
		return matches[2]
	}
	return ""
}

func getCalculator(input string) string {
	patt := "^(?i)Hasil dari (.+)|([0-9]+\\s*[-+*/%^]\\s*[0-9]+.*)$"
	reg := regexp.MustCompile(patt)
	matches := reg.FindStringSubmatch(input)
	if len(matches) == 3 {
		if len(matches[1]) > 0 {
			if matches[1][len(matches[1])-1] == '?' {
				return matches[1][:len(matches[1])-1]
			} else {
				return matches[1]
			}
		} else {
			if matches[2][len(matches[2])-1] == '?' {
				return matches[2][:len(matches[2])-1]
			} else {
				return matches[2]
			}
		}
	}
	return ""
}

func extractQuestionAnswer(input string) (string, string) {
	pattern := regexp.MustCompile(`(?i)^tambahkan pertanyaan\s(.+)\sdengan jawaban\s(.+)$`)
	matches := pattern.FindStringSubmatch(input)
	if len(matches) < 3 {
		return "", ""
	}
	question := strings.TrimSpace(matches[1])
	answer := strings.TrimSpace(matches[2])
	return question, answer
}

func getQuestionDeleteCommand(input string) string {
	reg := regexp.MustCompile(`(?i)hapus pertanyaan (.*)`)
	matches := reg.FindStringSubmatch(input)
	if len(matches) < 2 {
		return ""
	}
	question := strings.TrimSpace(matches[1])
	return question
}
