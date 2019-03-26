//Task 4.1 and 4.2 implementation from book page 112
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sha384flag = flag.Bool("sha384", false, "print value in sha384")
var sha512flag = flag.Bool("sha512", false, "print value in sha512")

func main() {
	flag.Parse()
	if len(os.Args) < 3 {
		fmt.Print("error:wrong number of arguments")
		return
	}
	if *sha384flag {
		if len(os.Args[2:]) != 2 {
			fmt.Print("error:wrong number of arguments")
			return
		}
		c1 := sha512.Sum384([]byte(os.Args[1]))
		c2 := sha512.Sum384([]byte(os.Args[2]))
		fmt.Printf("%x\n%x\n%08[1]b\n%08b\n", c1, c2)
	}
	if *sha512flag {
		if len(os.Args[2:]) != 2 {
			fmt.Print("error:wrong number of arguments")
			return
		}
		c1 := sha512.Sum512([]byte(os.Args[1]))
		c2 := sha512.Sum512([]byte(os.Args[2]))
		fmt.Printf("%x\n%x\n%[1]08b\n%08b\n", c1, c2)
	}
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("%x\n%x\n%t\n%T\n%08b\n%08b\n%08b\n", c1, c2, c1 == c2, c1, c1[0], c2[0], c1[0]^c2[0])
	fmt.Printf("%08b\n%08b\n%v", c1, c2, countAllDifBits(c1, c2))
}

func countAllDifBits(byteArr1, byteArr2 [32]byte) int {
	var globalCounter int
	for i, singleBit1 := range byteArr1 {
		if singleBit1 != byteArr2[i] {
			globalCounter += PopCountLoopDropBits(singleBit1, byteArr2[i])
		}
	}
	return globalCounter
}

// PopCountLoopDropBits returns degree of filling
// (number of bits set) value x using dropping right most non-zero bit
func PopCountLoopDropBits(x1, x2 byte) int {
	x := x1 ^ x2
	var counter int
	for x != 0 {
		counter++
		x &= (x - 1)
	}
	return counter
}
