// Упражнение 3.4.
// Следуя подходу, использованному в примере с фигурами Лиссажу из раздела 1.7, создайте веб-сервер, который вычисляет
// поверхности и возвращает клиенту SVG-данные. Сервер должен использовать в ответе заголовок ContentType наподобие
// следующего:
// w.Header().Set("ContentType", "image/svg+xml")
// (Этот шаг не был нужен в примере с фигурами Лиссажу, так как сервер использует стандартную эвристику распознавания
// распространенных форматов наподобие PNG по первым 512 байтам ответа и генерирует корректный заголовок.) Позвольте
// клиенту указывать разные параметры, такие как высота, ширина и цвет, в запросе HTTP.

//!+

// Surface вычисляет SVG-представление трехмерного графика функции.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
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
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer) {
	var s string
	s = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	out.Write([]byte(s))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			out.Write([]byte(s))
		}
	}
	s = fmt.Sprintln("</svg>")
	out.Write([]byte(s))
}

func corner(i, j int) (float64, float64) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Вычисляем высоту поверхности z.
	z := f(x, y)

	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
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
