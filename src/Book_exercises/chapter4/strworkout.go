// Task 4.5 (getRidOfDup) realization from book page 122
package main

import "fmt"

// nonempty returns cut, which contains only nonempty strings.
// The contents of the base array change when the function is running.
func nonempty(lines []string) []string {
	i := 0
	for _, s := range lines {
		if s != "" {
			lines[i] = s
			i++
		}
	}
	return lines[:i]
}

func nonempty2(lines []string) []string {
	out := lines[:0]
	for _, s := range lines {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func getRidOfDup(lines []string) []string {
	i := 0
	arrLen := len(lines)
	for j := 0; j < arrLen-1; j++ {
		if lines[j] != lines[j+1] {
			lines[i] = lines[j]
			i++
			if j == arrLen-2 {
				lines[i+1] = lines[j+1]
				i++
			}
		}
	}
	return lines[:i]
}

func main() {
	lines := []string{
		"",
		"gsagasg",
		"",
		"",
		"gshawaw",
		"yryrttset",
		""}
	fmt.Printf("%q\n", lines)
	// assignment is mandatory here
	lines = nonempty(lines)
	fmt.Printf("%q\n", lines)
	lines2 := []string{
		"",
		"gsagasg",
		"",
		"",
		"gaga",
		"gaga",
		"gaga",
		"",
		"sfasrar",
		"yryrttset",
		""}
	fmt.Printf("%q\n", lines2)
	// assignment is mandatory here
	lines2 = getRidOfDup(lines2)
	fmt.Printf("%q\n", lines2)

}
