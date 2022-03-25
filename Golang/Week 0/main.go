package main

import (
	"fmt"
	"strconv"
)

func printSerie(num int) string {
	if (num%3 == 0) && (num%5 == 0) {
		return "fizzbuzz"
	}
	if (num %3 == 0) {
		return "buzz"
	} 
	if (num % 5 == 0) {
		return "fizz"
	}
		return strconv.Itoa(num)
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(printSerie(i))
	}
}