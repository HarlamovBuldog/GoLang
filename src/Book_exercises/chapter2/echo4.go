// Echo4 prints command line arguments
// Example string:
// .\echo4 -n -s / pa ra bc jk
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "разделитель")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
