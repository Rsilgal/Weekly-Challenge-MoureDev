package main

import (
	"fmt"
)

func main() {
	var text string = "Hola caracola"

	fmt.Print(reverseString(text))
}

func reverseString(text string) string {
	var aux string = ""
	length := len(text) - 1

	for i := length; i >= 0; i-- {
		aux += fmt.Sprintf("%c", text[i])
	}

	return aux
}
