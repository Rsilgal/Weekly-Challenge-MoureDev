package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for thread := 1; thread <= 100; thread++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			for i := 2; i <= n; i++ {
				if (n%i == 0) && (i != n) {
					return
				}
			}
			fmt.Printf("El numero %d es un numero primo.\n", n)

		}(thread)
	}
	wg.Wait()
}
