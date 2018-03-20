// Упражнение 2.5.
// Выражение х&(х-1) сбрасывает крайний справа ненулевой бит х. Напишите версию PopCount, которая подсчитывает биты с
// использованием этого факта, и оцените ее производительность.

package popcount

import (
	"testing"
)

// -- Benchmarks --

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkPopCount(b *testing.B) {
	bench(b, PopCount)
}

func BenchmarkPopCountLoop(b *testing.B) {
	bench(b, PopCountLoop)
}

func BenchmarkPopCountByShifting(b *testing.B) {
	bench(b, PopCountByShifting)
}

func BenchmarkPopCountByClearing(b *testing.B) {
	bench(b, PopCountByClearing)
}

// goos: darwin
// goarch: amd64
// pkg: go/TheGoProgrammingLanguage/ch2/ex2.5
// 100000000	        10.1 ns/op
// 50000000	        31.2 ns/op
// 10000000	       237 ns/op
// 100000000	        27.5 ns/op
// PASS
