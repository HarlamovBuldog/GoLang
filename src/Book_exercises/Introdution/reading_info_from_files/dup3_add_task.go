//Programm reads standart input or list of named files.
//If this is standart input programm simply prints lines
//which appears more then 1 time.
//If these are files then program counts how many times
//each string appears in program for each file and also
//says if there are any repeated strings.
//Note: Task 1.4 realization from book page 34
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
			doLinesRepeat := false
			countsForEachFile := make(map[string]int)

			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, arg+": %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()

			f1, err := os.Open(arg)
			fmt.Println(arg + ":")
			if err != nil {
				fmt.Fprintf(os.Stderr, arg+": %v\n", err)
				continue
			}
			countLines(f1, countsForEachFile)
			f1.Close()

			for line, n := range countsForEachFile {
				fmt.Printf("%d\t%s\n", n, line)
				if doLinesRepeat != true && n > 1 {
					doLinesRepeat = true
				}
			}

			if doLinesRepeat {
				fmt.Println("File " + arg + " has repeated lines")
			} else {
				fmt.Println("File " + arg + " doesn't have repeated lines")
			}
		}
	}

	fmt.Println("Total:")
	for line, n := range counts {
		if n > 1 {
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
