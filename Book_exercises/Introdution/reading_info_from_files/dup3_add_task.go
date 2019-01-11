//Dup2 prints text of each line, which
//appears in input more then 1 times.
//Programm reads standart input or list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			countsForEachFile := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, arg+": %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()

			f1, err1 := os.Open(arg)
			fmt.Println(arg + ":")
			if err1 != nil {
				fmt.Fprintf(os.Stderr, arg+": %v\n", err1)
				continue
			}
			countLines(f1, countsForEachFile)
			f1.Close()

			for line, n := range countsForEachFile {
				if len(countsForEachFile) > 1 {
					fmt.Printf("%d\t%s\n", n, line)
				}
			}
		}
	}

	fmt.Println("Total:")
	for line, n := range counts {
		if len(counts) > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//Note: Ignoring potential errors
	//from input.Err()
}
