package main

import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(x int) {
			println(x)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
