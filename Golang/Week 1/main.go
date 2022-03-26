package main

import "fmt"

func isAnagram(wordA, wordB string) bool {

	// var m map[string]int

	if wordA == wordB {
		return false
	}

	return false
}

func countOfChar(word string) map[string]int {
	aux := make(map[string]int)
	num := 1

	for _, character := range word {
		char := fmt.Sprintf("%c", character)
		aux[char] += num
	}

	return aux
}

/*
* 	Nepal - Panel
*	Raza - Zara
*	Mora - Roma
 */

func main() {
	
}
