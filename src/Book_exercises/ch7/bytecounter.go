package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int
type WordsCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// Exercise 7.1 from book realization
func (c *WordsCounter) Write(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	counter := 0
	for scanner.Scan() {
		counter++
	}
	if err := scanner.Err(); err != nil {
		return counter, err
	}
	*c += WordsCounter(counter)
	return counter, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) //"5", = len("hello")
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	var cw WordsCounter
	cw.Write("hi there u beaty")
	fmt.Println(cw) // "4"
}
