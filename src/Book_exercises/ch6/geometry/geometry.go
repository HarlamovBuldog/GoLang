package geometry

import (
	"image/color"
	"math"
)

// Point is the exported type with exported fields
type Point struct{ X, Y float64 }

// Path - path of points, connected by straight line segments
// Path is just another type to show that we can have one
// more valid name Distance for method with another GETTER type
type Path []Point

// ColoredPoint is struct in which Point is build-in
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Distance is the raditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance this time is the same as previous,
// but stands out as method of the type Point
// p Point is GETTER in this case
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy multiplies both Point fields by given factor
// ScaleBy uses pointer, so original p *Point will be changed
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Distance returns path length
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
