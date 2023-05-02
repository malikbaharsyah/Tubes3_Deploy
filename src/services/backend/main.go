package main

import (
	"example.com/algorithm"
	"fmt"
)

func main() {
	// fmt.Println("Hello World!")
	fmt.Println(algorithm.BoyerMooreAlgorithm("This is A TEST test", "is"))
	fmt.Println(algorithm.KnuthMorrisPratt("This is A TEST test", "test"))
	algorithm.Calculator("24+6/2*10")
	algorithm.Calendar("29/2/2016")
}
