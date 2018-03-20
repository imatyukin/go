// Упражнение 2.4.
// Напишите версию PopCount, которая подсчитывает биты с помощью сдвига аргумента по всем 64 позициям, проверяя при
// каждом сдвиге крайний справа бит. Сравните производительность этой версии с выборкой из таблицы.

package popcount

import (
	"testing"
)

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

// goos: darwin
// goarch: amd64
// pkg: go/TheGoProgrammingLanguage/ch2/ex2.4
// 2000000000	         0.46 ns/op
// 50000000	        30.5 ns/op
// 5000000	       390 ns/op
// PASS
