package main

import (
	"fmt"
	"sync"
)

func main() {
	// reduce this number and see the effect
	const workers = 99

	var wg sync.WaitGroup
	wg.Add(workers)
	m := map[int]int{}
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				m[i]++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}
