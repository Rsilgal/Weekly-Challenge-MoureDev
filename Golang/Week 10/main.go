package main

import (
	"fmt"
)

func countExpresions(text string) map[string]int {
	result := make(map[string]int)

	for _, element := range text {
		character := fmt.Sprintf("%c", element)
		switch character {
		case "{", "}", "(", ")", "[", "]":
			result[character] += 1
		}
	}
	return result
}

func checkBalance(countDic map[string]int) bool {
	a := countDic["{"] == countDic["}"]
	b := countDic["("] == countDic[")"]
	c := countDic["["] == countDic["]"]

	return a && b && c
}

func main() {
	expression := "{ [ a * ( c + d ) ] - 5 }"
	counterExpresion := countExpresions(expression)
	fmt.Printf("Test result: %v", checkBalance(counterExpresion))
}
