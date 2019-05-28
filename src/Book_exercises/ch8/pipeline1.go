// Channel usage example from book
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
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squaring
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Print (in main go-subprogram)
	for {
		fmt.Println(<-squares)
		time.Sleep(1 * time.Second)
	}
}
