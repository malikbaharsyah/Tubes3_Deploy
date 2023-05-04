package algorithm

import (
	"fmt"
	"regexp"
)

func ParseInput(input string, listOfQuestion []string, listOfAnswer []string) string {
	patt1 := "^(?i)(Apakah|Apa|Bagaimana|Kenapa|Siapa|Mengapa|Kapan|Di mana)\\s.+\\??"
	reg1 := regexp.MustCompile(patt1)

	patt2 := "^(?i)(Hari\\s\\d{1,2}/\\d{1,2}/\\d{4}|\\d{1,2}/\\d{1,2}/\\d{4})\\??$"
	reg2 := regexp.MustCompile(patt2)

	patt3 := `^(Hasil dari\s*|Hasil\s*)?-?\s*(\(\s*)*\s*-?\s*\d+(\.\d+)?(\s*\))*\s*(\s*[-+*/]\s*(\(\s*)*\s*-?\s*(\(\s*)*\d+(\.\d+)?(\s*\))*\s*)*\s*(\?\s*|\s\?\s*)?$`
	reg3 := regexp.MustCompile(patt3)

	if reg1.MatchString(input) {
		fmt.Println("Question")
		input = getQuestion(input)
		fmt.Println(input)
		index, persen := searchQuestion(input, listOfQuestion)
		if persen > 0.9 {
			return listOfAnswer[index]
		} else {
			return "Invalid input"
		}
	} else if reg2.MatchString(input) {
		fmt.Println("Date")
		input = getDate(input)
		fmt.Println(input)
		return Calendar(input)
	} else if reg3.MatchString(input) {
		fmt.Println("Calculator")
		input = getCalculator(input)
		fmt.Println(input)
		return Calculator(input)
	} else {
		return "Invalid input"
	}
}

func getQuestion(input string) string {
	patt := "^(?i)(Apakah|Apa|Bagaimana|Kenapa|Siapa|Mengapa|Kapan|Di mana)\\s(.+)\\?$"
	reg := regexp.MustCompile(patt)
	matches := reg.FindStringSubmatch(input)
	if len(matches) == 3 {
		return matches[2]
	}
	return ""
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
