package main

import (
	"fmt"
	"strings"
)

func deleteSpaces(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func formatText(text string) string {
	_text := deleteSpaces(text)
	_text = strings.ToLower(_text)
	return _text
}

func isPalindrome(text string) bool {
	text = formatText(text)

	for i := 0; i < len(text); i++ {

		c := fmt.Sprintf("%c", text[i])
		c2 := fmt.Sprintf("%c", text[len(text)-1-i])

		if c != c2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Print(isPalindrome("Ana lleva al oso la avellana"))
}
