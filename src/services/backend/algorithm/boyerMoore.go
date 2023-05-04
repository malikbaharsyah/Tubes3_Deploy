package algorithm

func BoyerMooreAlgorithm(text, pattern string) int {
	badCharTable := [256]int{}
	for i := range badCharTable {
		badCharTable[i] = len(pattern)
	}
	for i := 0; i < len(pattern); i++ {
		badCharTable[pattern[i]] = max(1, len(pattern)-i-1)
	}

	i := len(pattern) - 1
	for i < len(text) {
		j := len(pattern) - 1
		for j >= 0 && pattern[j] == text[i] {
			i--
			j--
		}
		if j < 0 {
			return i + 1
		}
		i += badCharTable[text[i]]
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}