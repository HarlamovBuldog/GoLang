package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "https://gopl.io"
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Server doesn't work: %v\n", err)
	}
	// or
	// log customization
	//log.SetPrefix("wait: ")
	//log.SetFlags(0)
	if err1 := WaitForServer(url); err1 != nil {
		log.Fatalf("Server doesn't work: %v\n", err1)
	}
}

// WaitForServer is trying to connect to server with given URL.
// Tries are executed for 1 minute with growing intervals.
// Logs error, if all tries are unlucky.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil // Successful connection
		}
		log.Printf("Server is not responding (%s); repeat...", err)
		time.Sleep(time.Second << uint(tries)) // Increasing delay
	}
	return fmt.Errorf("Server %s is not responding; time %s ", url, timeout)
}
