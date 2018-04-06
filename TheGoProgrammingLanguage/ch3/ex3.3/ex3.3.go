// Упражнение 3.3.
// Окрасьте каждый многоугольник цветом, зависящим от его высоты, так, чтобы пики были красными (#ff0000), а низины —
// синими (#0000ff).

//!+

// Surface вычисляет SVG-представление трехмерного графика функции.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // Размер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапазон осей (-xyrange..+ xyrange)
	xyscale       = width / 2 / xyrange // Пикселей в единице х или у
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // Углы осей х, у (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var isPeak bool
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, isPeak := corner(i+1, j+1)
			var color string
			if isPeak {
				color = "#ff0000"
			} else {
				color = "#0000ff"
			}
			fmt.Printf("<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Вычисляем высоту поверхности z.
	z := f(x, y)

	var b bool
	if z >= 0 {
		b = true
	} else {
		b = false
	}

	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, b
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0,0)
	if r != 0 {
		return math.Sin(r) / r
	} else {
		return 0
	}
}

//!-
