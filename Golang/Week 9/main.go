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

	text = ".|.---"

	fmt.Println(MorseCodeToAlpha(text, morseDic))
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

func MorseCodeToAlpha(textMorseCode string, morseDic map[string]string) string {
	var alpha string = ""

	characterList := strings.Split(textMorseCode, "|")

	for _, e := range characterList {
		alpha += getKey(e,morseDic)
	}

	return alpha
}

func getKey(morseCode string, morseDic map[string]string) string {
	for i, elem := range morseDic {
		if elem == morseCode {
			return i
		}
	}
	return ""
}
