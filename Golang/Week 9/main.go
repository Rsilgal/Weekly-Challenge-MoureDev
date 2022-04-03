package main

import (
	"fmt"
	"morse/morsecode"
	"strings"
)

func main() {
	morseDic := morsecode.InitMorseCode()

	var text string = "Hola Caracola"
	var morseTranslate string = ""

	text = strings.ToLower(text)

	for _, character := range text {
		aux := fmt.Sprintf("%c", character)
		morseTranslate += morseDic[aux]
	}

	fmt.Println(morseTranslate)
}
