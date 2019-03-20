//Task 2.3 realization from book page 69
package main

import (
	"fmt"
	"os"
	"strconv"
)

// pc[i] - number of single bits in i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		value, err := strconv.ParseUint(arg, 0, 64)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%08b %v\n", value, PopCountLoop(value))
		}
	}
	var x uint64 = 2572572752525252525
	var y uint64 = 56

	fmt.Printf("%08b %08b %08b\n", x, y, x>>y)
}

// PopCountLoop returns degree of filling
// (number of bits set) value x using loop
func PopCountLoop(x uint64) int {
	var byteArray byte
	for i := uint64(0); i < 8; i++ {
		byteArray += pc[byte(x>>(i*8))]
	}
	return int(byteArray)
}

// PopCountOneExpr returns degree of filling
// (number of bits set) value x using one expression
func PopCountOneExpr(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
