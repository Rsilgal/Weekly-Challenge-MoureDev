package main

import "fmt"

/*
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
*/

func calcFib(numOfElements int) {
	f1 := 0
	f2 := 1

	for i := 1; i <= numOfElements; i++ {
		printFibonacci(f1, i == numOfElements)
		f1, f2 = f2, (f1 + f2)
	}
}

func printFibonacci(num int, isLastElem bool) {
	if isLastElem {
		fmt.Printf("%d", num)
		return
	}
	fmt.Printf("%d, ", num)
}

func main() {
	calcFib(50)
}
