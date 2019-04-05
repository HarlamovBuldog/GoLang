// recursive tree sort
// represatation of the fact that structure
// variable can be type of the same structure
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values "on place"
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues adds elements t to values in needed order
// and returns final slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// The same as to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	input := []int{5, 2, 8, 0, 9, 4, 2, 1, 3, 4}
	fmt.Printf("%v\n", input)
	Sort(input)
	fmt.Printf("%v\n", input)
}
