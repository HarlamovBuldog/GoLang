// Task 3.10 (commaNonRecur)
// Task 3.11 (commaNonRecur)
// implementation from book page 101
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// comma inserts commas into string representation
// of a non-negative decimal number using recursive algorithm
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// commaNonRecur inserts commas into string representation
// of a non-negative decimal number using
// non-recursive algorithm and bytes.Buffer
// Note:seems like this is workable but ugly way of doing things
// Looking for better solution
func commaNonRecur(s string) string {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err.Error()
	}
	var prefixSign = ""
	var floatSuffix = ""
	if signIndex := strings.IndexAny(s, "+-"); signIndex != -1 {
		prefixSign = s[:signIndex+1]
		s = s[signIndex+1:]
	}
	if dotIndex := strings.IndexAny(s, "."); dotIndex != -1 {
		floatSuffix = s[dotIndex:]
		s = s[:dotIndex]
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, j := n, 0; i > 0; i, j = i-3, j+1 {
		tempStr := buf.String()
		buf.Reset()
		if i >= 3 {
			buf.WriteString(s[i-3:])
			if j != 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(tempStr)
			s = s[:i-3]
		} else {
			buf.WriteString(s)
			buf.WriteByte(',')
			buf.WriteString(tempStr)
		}
	}
	tempStr := buf.String()
	buf.Reset()
	buf.WriteString(prefixSign)
	buf.WriteString(tempStr)
	buf.WriteString(floatSuffix)
	return buf.String()
}

// intsToString is similar to fmt.Sprint(values), but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	for _, strArg := range os.Args[1:] {
		var intSlice []int
		intArg, err := strconv.ParseInt(strArg, 0, 64)
		if err != nil {
			fmt.Printf("%s => %s => %s\n", strArg, commaNonRecur(strArg), "int: failed to parse")
			continue
		}
		for intArg != 0 {
			intSlice = append(intSlice, int(intArg%10))
			intArg /= 10
		}
		reverse(intSlice)
		fmt.Printf("%s => %s => %s\n", strArg, commaNonRecur(strArg), intsToString(intSlice))
	}
}

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}
