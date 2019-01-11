//dup1 выводит текст каждой строки, которая появляется в
//стандарном вводе более одного раза, а также количество
//ее появлений
//Note: Ctrl+Z in console for end entering
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	//MAP is pair called "key-value"
	//KEY TYPE is string
	//VALUE TYPE is int
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "stop" {
			break
		}
		counts[input.Text()]++
		//shorting for following code:
		//line := input.Text()
		//counts[line] = counts[line] + 1
	}
	//Note: Ignoring potential errors
	//from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
