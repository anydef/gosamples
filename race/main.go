package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 6; i++ {
		fmt.Println(i) // Not the 'i' you are looking for.
		go func(i int) {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}(i)
	}
	wg.Wait()
}
