package main

import (
	"fmt"
	"strconv"
)

func DecToBin(number int) string {
	var binaryNumber []int
	auxNumber := number

	binaryNumber = calcBinary(auxNumber, binaryNumber)

	resultString := revereSlice(binaryNumber)

	return resultString
}

func calcBinary(auxNumber int, binaryNumber []int) []int {
	for auxNumber >= 1 {
		if auxNumber%2 == 0 {
			binaryNumber = append(binaryNumber, 0)
			auxNumber = auxNumber / 2
			continue
		}
		binaryNumber = append(binaryNumber, 1)
		auxNumber = auxNumber / 2
	}
	return binaryNumber
}

func revereSlice(slice []int) string {
	result := ""

	for i := len(slice) - 1; i >= 0; i-- {
		result += strconv.Itoa(slice[i])
	}

	return result
}

func main() {
	fmt.Print(DecToBin(456))
}
