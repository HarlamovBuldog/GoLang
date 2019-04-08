//Task 2.3, 2.4 and 2.5 realization from book page 69
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
			//fmt.Printf("%08b %v\n", value, PopCountLoop(value))
			//fmt.Printf("%08b %v\n", value, PopCountOneExpr(value))
			//fmt.Printf("%08b %v\n", value, PopCountLoop64Pos(value))
			fmt.Printf("%08b %v\n", value, PopCountLoopDropBits(value))
		}
	}
}

// PopCountLoop returns degree of filling
// (number of bits set) value x using loop and pc byte table
func PopCountLoop(x uint64) int {
	var byteArray byte
	for i := uint64(0); i < 8; i++ {
		byteArray += pc[byte(x>>(i*8))]
	}
	return int(byteArray)
}

// PopCountOneExpr returns degree of filling
// (number of bits set) value x using one expression and pc byte table
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

// PopCountLoop64Pos returns degree of filling
// (number of bits set) value x using one loop for each of 64 positions
func PopCountLoop64Pos(x uint64) int {
	var counter byte
	for i := uint64(0); i < 64; i++ {
		counter += byte((x >> i) & 1)
	}
	return int(counter)
}

// PopCountLoopDropBits returns degree of filling
// (number of bits set) value x using dropping right most non-zero bit
func PopCountLoopDropBits(x uint64) int {
	var counter byte
	for x != 0 {
		counter++
		x &= (x - 1)
	}
	return int(counter)
}
