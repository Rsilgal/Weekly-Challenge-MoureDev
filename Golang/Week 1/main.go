package main

import (
	"fmt"
	"reflect"
	"strings"
)

func isAnagram(wordA, wordB string) bool {

	_wordA := strings.ToLower(wordA)
	_wordB := strings.ToLower(wordB)

	if _wordA == _wordB {
		return false
	}

	return reflect.DeepEqual(countOfChar(_wordA), countOfChar(_wordB))
}

func countOfChar(word string) map[string]int {
	aux := make(map[string]int)

	for _, character := range word {
		char := fmt.Sprintf("%c", character)
		aux[char] += 1
	}

	return aux
}

/*
* 	Nepal - Panel
*	Raza - Zara
*	Mora - Roma
 */

func main() {
	fmt.Print(isAnagram("roma", "Moras"))
}
