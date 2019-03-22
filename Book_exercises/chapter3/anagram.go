// Task 3.12 (areAnagrams)
// implementation from book page 101
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func areAnagrams(s1, s2 string) bool {
	returnValue := true
	if utf8.RuneCountInString(s1) != utf8.RuneCountInString(s2) {
		return false
	}
	for _, rune2 := range s2 {
		if strings.ContainsRune(s1, rune2) == false ||
			countRune(s1, rune2) != countRune(s2, rune2) {
			returnValue = false
			break
		}
	}
	return returnValue
}

func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}

func main() {
	s1 := "tghyubnf"
	s2 := "ybfntguh"
	fmt.Printf("%s %s %v", s1, s2, areAnagrams(s1, s2))
}
