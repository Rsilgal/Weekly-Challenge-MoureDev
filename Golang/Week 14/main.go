package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func chekArmstrong(num int) bool {
	var total int = 0
	numToString := strconv.Itoa(num)

	for _, elem := range numToString {
		elemStr := fmt.Sprintf("%c", elem)
		a, err := strconv.Atoi(elemStr)

		if err != nil {
			log.Fatal(err.Error())
		}

		pow := math.Pow(float64(a), float64(len([]rune(numToString))))
		total += int(pow)
	}
	return total == num
}

func main() {
	var num int = 407

	fmt.Print(chekArmstrong(num))
}
