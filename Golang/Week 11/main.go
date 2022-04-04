package main

import (
	"fmt"
	"strings"
)

func existCharacter(str, char string) bool {
	return strings.Index(str, char) > -1
}



func deleteCharacters(str1, str2 string) (out string) {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	for _, elem := range str1 {
		aux := fmt.Sprintf("%c", elem)
		if existCharacter(str2, aux) == false {
			out += aux
		}
	}

	return out
}

func main() {
	a := "abcd"
	b := "Bf"
	c := deleteCharacters(a, b)
	d := deleteCharacters(b,a)
	fmt.Printf("Result First String: %s\nResult Second String: %s", c, d)
}
