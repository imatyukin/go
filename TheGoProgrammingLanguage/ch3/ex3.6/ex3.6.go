// Упражнение 3.6.
// Супервыборка (supersampling) — это способ уменьшить эффект пикселизации путем вычисления значений цвета в нескольких
// точках в пределах каждого пикселя и их усреднения. Проще всего разделить каждый пиксель на четыре “подпикселя”.
// Реализуйте описанный метод.

//!+

// Mandelbrot создает PNG-изображение фрактала Мандельброта.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	width, height := 1024, 1024

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	width = width * 2
	height = height * 2
	for py := 0; py < height; py = py + 2 {
		y1 := float64(py)/float64(height)*(ymax-ymin) + ymin
		y2 := float64(py+1)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px = px + 2 {
			x1 := float64(px)/float64(width)*(xmax-xmin) + xmin
			x2 := float64(px+1)/float64(width)*(xmax-xmin) + xmin
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)
			// Точка (px, py) представляет комплексное значение z.
			img.Set(px/2, py/2, color.Gray{(mandelbrot(z1).Y + mandelbrot(z2).Y + mandelbrot(z3).Y + mandelbrot(z4).Y) / 4.0})
		}
	}
	png.Encode(os.Stdout, img) // Примечание: игнорируем ошибки
}

func mandelbrot(z complex128) color.Gray {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Gray{0}
}

//!-

// Некоторые другие интересные функции:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
