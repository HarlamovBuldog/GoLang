// Task 4.9 realization from book page 129
// program counts word frequency using input file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	wordfreq := make(map[string]int)
	for _, file := range files {
		filePtr, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, file+": %v\n", err)
			continue
		}
		input := bufio.NewScanner(filePtr)
		//we are using split function bufio.ScanWords
		//to split text into words, but not into strings as usually
		input.Split(bufio.ScanWords)
		for input.Scan() {
			wordfreq[input.Text()]++
		}
		filePtr.Close()
	}
	fmt.Fprint(os.Stdout, "word\t\tcount\n")
	for word, count := range wordfreq {
		fmt.Fprintf(os.Stdout, "%s\t\t%d\n", word, count)
	}
}
