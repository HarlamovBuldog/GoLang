//dup1 выводит текст каждой строки, которая появляется в
//стандарном вводе более одного раза, а также количество
//ее появлений
//Note: Ctrl+Z in console for end entering
//Note: or just much easier way to specify any string
//that will break for cycle
//In our example it's string 'stop'
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//MAP is pair called "key-value"
	//KEY TYPE is string
	//VALUE TYPE is int
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "stop" {
			break
		}
		//shortcut for the following code:
		//line := input.Text()
		//counts[line] = counts[line] + 1
		counts[input.Text()]++

	}
	//Note: Ignoring potential errors
	//from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
