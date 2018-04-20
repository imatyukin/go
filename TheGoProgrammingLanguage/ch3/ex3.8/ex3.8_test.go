// Упражнение 3.8.
// Визуализация фракталов при высоком разрешении требует высокой арифметической точности. Реализуйте один и тот же
// фрактал с помощью четырех различных представлений чисел: complex64, complexl28, big.Float и big.Rat. (Два последних
// типа вы найдете в пакете math/big. Float использует произвольную, но ограниченную точность для чисел с плавающей
// точкой; Rat обеспечивает неограниченную точность для рациональных чисел.) Сравните производительность и потребление
// памяти при использовании разных типов. При каком уровне масштабирования артефакты визуализации становятся видимыми?

package main

import (
	"image/color"
	"testing"
)

func benchmarkMandelbrot(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotComplex128(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrot)
}

func BenchmarkMandelbrotComplex64(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrot64)
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrotBigFloat)
}

func BenchmarkMandelbrotRat(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrotRat)
}
