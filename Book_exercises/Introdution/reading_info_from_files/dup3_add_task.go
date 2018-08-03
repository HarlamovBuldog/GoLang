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
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			for _, arg := range files {
				f, err := os.Open(arg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
					continue
				}

				f.Close()
			}
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

func countLinesForEachFile(fileList range, )