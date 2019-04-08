// Task 4.8 realization from book page 129
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func getTypeOfCh(r rune) string {
	runeType := "undef"
	if unicode.IsDigit(r) {
		runeType = "Digit"
	} else if unicode.IsLetter(r) {
		runeType = "Letter"
	} else if unicode.IsPunct(r) {
		runeType = "Punct ch"
	} else if unicode.IsSpace(r) {
		runeType = "Space"
	} else if unicode.IsSymbol(r) {
		runeType = "Symbol"
	}
	return runeType
}

func main() {
	counts := make(map[rune]int)
	countsTypes := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		countsTypes[getTypeOfCh(r)]++
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d wrong symbols UTF-8\n", invalid)
	}
	fmt.Print("\ntype\tcount\n")
	for rType, n := range countsTypes {
		fmt.Printf("%v\t%d\n", rType, n)
	}
}
