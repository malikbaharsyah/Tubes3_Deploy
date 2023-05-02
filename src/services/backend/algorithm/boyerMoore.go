package algorithm

func BoyerMooreAlgorithm(text, pattern string) int {
 // fmt.Println(text, pattern)
 // fmt.Println(text[0], ", ", text[1], ", ", text[2, ", ", text[3], ", ", text[4], ", ", text[5], ", ", text[6], ", ", text[7], ", ", text[8])
 // fmt.Println("t ", patten[0], ", e ", pattern[1], ", s ", pattern[2], ", t ", pattern[3])
 badCharTable := [256]int{}
 for i := range badCharTable {
  badCharTable[i] = len(pattern)
  // fmt.Println(badCharTable[], i)
 }
 // fmt.Println(badCharTable)
 for i := 0; i < len(pattern); i++ {
  badCharTable[pattern[i]] = max(1, len(pattern)-i-1)
  // fmt.Println("Letters : ", patten[i], "  ->  Values : ", badCharTable[pattern[i]])
 }
 // fmt.Println(badCharTable)

 i := len(pattern) - 1
 for i < len(text) {
  j := len(pattern) - 1
  for j >= 0 && pattern[j] == text[i] {
   // fmt.Println(pattern[j], text[i])
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

// func main() {
//  fmt.Println(boyerMooreAlgorithm("This is A TEST test", "test"))
//  // fmt.Printn(boyerMooreAlgorithm("Apa yang sedang kamu lakukan disini!", "kamu"))
// }