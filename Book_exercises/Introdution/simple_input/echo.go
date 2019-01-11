//выводит аргументы
//коммандной строки
//реализация через цикл и один Println
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
