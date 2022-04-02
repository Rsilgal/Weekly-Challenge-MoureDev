package main

import (
	"fmt"
	"regexp"
	"strings"
)

func formatText(text string, re *regexp.Regexp) []string {
	return re.Split(text, -1)
}

func countDiferentsWords(textFormated []string) map[string]int {
	counter := make(map[string]int)

	for _, element := range textFormated {
		if element == "" {
			continue
		}
		elementLower := strings.ToLower(element)
		counter[elementLower] += 1
	}

	return counter
}

func buildRegex() *regexp.Regexp {
	return regexp.MustCompile(`(,\s)|(\.*\.\s)|(\s)|(\.)$`)
}

func main() {
	re := buildRegex()

	text := "Hola, esto Es es un texto bastante aleatorio."

	textFormated := formatText(text, re)

	mapOfWords := countDiferentsWords(textFormated)

	fmt.Print(mapOfWords)
}
