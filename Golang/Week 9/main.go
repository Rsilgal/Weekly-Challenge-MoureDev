package main

import (
	"fmt"
	"morse/morsecode"
	"strings"
)

func main() {
	morseDic := morsecode.InitMorseCode()

	var text string = "Hola Caracola"
	morseTranslate := AlphaToMorseCode(text, morseDic)

	fmt.Println(morseTranslate)
}

func AlphaToMorseCode(text string, morseDic map[string]string) string {
	var morseTranslate string = ""

	text = strings.ToLower(text)

	for _, character := range text {
		aux := fmt.Sprintf("%c", character)
		morseTranslate += morseDic[aux]
	}
	return morseTranslate
}
