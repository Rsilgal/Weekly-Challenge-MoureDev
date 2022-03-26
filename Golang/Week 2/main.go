package main

import "fmt"

func calcFibonacci(listInput []int, numOfElements int) []int {
	list := listInput
	for i := 1; i <= numOfElements-2; i++ {
		list = append(list, list[i-1]+list[i])
	}
	return list
}

func printSlice(list []int) {
	numOfElemnts := len(list) - 1
	for index, elem := range list {
		if index == numOfElemnts {
			fmt.Printf("%d", elem)
		} else {
			fmt.Printf("%d, ", elem)
		}
	}
}

func main() {
	var fib []int
	fib = append(fib, 0)
	fib = append(fib, 1)

	list := calcFibonacci(fib, 50)
	printSlice(list)
}

/*
0
1
1
2
3
5
8
13
21
34
55
89
...
*/
