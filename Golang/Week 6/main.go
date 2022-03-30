package main

import (
	"fmt"
)

func main() {
	var a string = "Hola caracola"

	for _, elem := range a {
		fmt.Println(fmt.Sprintf("%c", elem))

	}
}
