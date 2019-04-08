//выводит аргументы
//коммандной строки
//Note: Task 1.1 and Task 1.2 realization from book page 29
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + " " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}
