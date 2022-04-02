package main

import (
	"fmt"
	"regexp"
)

func formatText(text string, re *regexp.Regexp) []string {
	return re.Split(text, -1)
}

func countDiferentsWords() map[string]int {
	counter := make(map[string]int)

	
	return map[string]int{}
}

func buildRegex() *regexp.Regexp {
	return regexp.MustCompile(`(,\s)|(\.*\.\s)|(\s)|(\.)$`) 
}

func main() {
	re := buildRegex()

	text := "Hola, esto es un texto bastante aleatorio."

	a := formatText(text, re)

	for _, c := range a {
		fmt.Println(c)
	}
}
