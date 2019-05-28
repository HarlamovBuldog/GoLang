// Channel usage example from book with close(ch_name)
package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Creation
	go func() {
		for x := 0; x < 15; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squaring
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Print (in main go-subprogram)
	for x := range squares {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
