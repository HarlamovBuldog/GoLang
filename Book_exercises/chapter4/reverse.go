// Task 4.3 (reversePtr), 4.4(rotate) implementation from book page 122
package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	var b []int
	b = append(b, 6, 8, 9, 5, 4, 3, 2, 0, 1)
	fmt.Println(b)
	reversePtr(&b)
	fmt.Println(b)
	rotate(b[:], 5)
	fmt.Println(b)
	fmt.Println(a)
	rotate(a[:], 5)
	fmt.Print(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reversePtr(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func rotate(s []int, step int) {
	if step >= len(s) || step <= 0 {
		fmt.Println("Wrong step")
		return
	}
	for k := 0; k < step; k++ {
		for i := 0; i < len(s)-1; i++ {
			s[i], s[i+1] = s[i+1], s[i]
		}
	}
}
