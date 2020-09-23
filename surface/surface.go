//Learning URL https://books.studygolang.com/gopl-zh/
package surface

import (
	"fmt"
	"math"
	"strings"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func MakeSvg() string {
	svg := strings.Builder{}
	svgHeader := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	svg.WriteString(svgHeader)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, cc := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			pointContent := fmt.Sprintf("<polygon style='fill: #%x' points='%g,%g %g,%g %g,%g %g, %g'/>\n ",
				cc, ax, ay, bx, by, cx, cy, dx, dy)
			svg.WriteString(pointContent)
		}
	}
	svg.WriteString("</svg>")
	return svg.String()
}
func corner(i, j int) (float64, float64, int) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	color := int((0xff0000-0x0000ff)*(z+0.3)/1.3 + 0x0000ff)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	//sy := height/2 + (x+y)*sin30*xyscale
	return sx, sy, color
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sinh(r) / r
}
func f1(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
func f2(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Tanh(r) / r
}
