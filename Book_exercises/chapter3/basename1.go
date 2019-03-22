// basename deletes file path components and suffix of file type.
package main

import (
	"fmt"
	"os"
	"strings"
)

func basenameLibFunc(s string) string {
	slash := strings.LastIndex(s, "/") // -1, if "/" is not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func basename(s string) string {
	// drop last symbol '/' and everything before it
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// save everything before last dot '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	for _, strArg := range os.Args[1:] {
		fmt.Printf("%s => %s => %s\n", strArg, basename(strArg), basenameLibFunc(strArg))
	}
}
