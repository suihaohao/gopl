//Learning URL https://books.studygolang.com/gopl-zh/
package geometry

import (
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Add(q Point) Point {
	return Point{p.X+q.X, p.Y+q.Y}
}
func (p Point) Sub(q Point) Point {
	return Point{p.X-q.X, p.Y-q.Y}
}

func (path Path) TranslateBy(offset Point, add bool)  {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	}else {
		op= Point.Sub
	}

	for i := range path{
		path[i] = op(path[i], offset)
	}
}