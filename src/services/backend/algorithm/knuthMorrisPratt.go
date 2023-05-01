/*
Knuth-Morris-Pratt algorithm for pattern matching
*/

package main

import "fmt"

func knuthMorrisPratt(text, pattern string) int {
	m, i := 0, 0
	table := make([]int, len(text))
	fmt.Println(table)
	kmpTab(pattern, table)
	for m+i < len(text) {
		if pattern[i] == text[m+i] {
			if i == len(pattern)-1 {
				return m
			}
			i++
		} else {
			fmt.Println("Table sebelumnya: ", table)
			if table[i] > -1 {
				i = table[i]
				m = m + i - table[i]
				fmt.Println("Table sekarang  : ", table)
			} else {
				i = 0
				m++
			}
		}
	}
	return len(text)
}

func kmpTab(pattern string, table []int) {
	position, candidate := 2, 0
	table[0], table[1] = -1, 0
	for position < len(pattern) {
		if pattern[position-1] == pattern[candidate] {
			candidate++
			table[position] = candidate
			position++
		} else if candidate > 0 {
			candidate = table[candidate]
		} else {
			table[position] = 0
			position++
		}
	}
}
