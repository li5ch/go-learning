package main

import "fmt"

func main() {
	for i := 17; i <= 21; i++ { // write
		go func() { /* Create a new goroutine */

			fmt.Println(i) // read

		}()

	}
	for {

	}
}
