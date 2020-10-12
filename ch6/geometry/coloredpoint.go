//Learning URL https://books.studygolang.com/gopl-zh/
package geometry

import (
	"image/color"
)
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p ColoredPoint) Distance(q Point) float64  {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64)  {
	p.Point.ScaleBy(factor)
}