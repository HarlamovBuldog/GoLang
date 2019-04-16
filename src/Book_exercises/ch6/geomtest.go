package main

import (
	"Book_exercises/ch6/geometry"
	"fmt"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q)) // "5", function call
	// p.Distance is SELECTOR
	fmt.Println(p.Distance(q)) // "5", method call

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"

	//> 3 ways to work with
	r := &geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"

	p1 := geometry.Point{3, 4}
	p1ptr := &p1
	p1ptr.ScaleBy(2)
	fmt.Println(p1) // "{6, 8}"

	p2 := geometry.Point{3, 3}
	(&p2).ScaleBy(3)
	fmt.Println(p2) // "{9, 9}"
	//<

	// Point is build-in field of structure ColoredPoint
	// So there is some syntactic sugare we can actually use
	var cp geometry.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	distanceFromP := p.Distance   // Value-method
	fmt.Println(distanceFromP(q)) // "5"

	distance := geometry.Point.Distance // Expression-method
	fmt.Println(distance(p, q))         // "5"
}
