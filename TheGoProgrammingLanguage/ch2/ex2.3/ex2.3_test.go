// Упражнение 2.3.
// Перепишите функцию PopCount так, чтобы она использовала цикл вместо единого выражения. Сравните производительность
// двух версий. (В разделе 11.4 показано, как правильно сравнивать производительность различных реализаций.)

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

// goos: darwin
// goarch: amd64
// pkg: go/TheGoProgrammingLanguage/ch2/ex2.3
// 2000000000	         0.48 ns/op
// 50000000	        30.7 ns/op
// PASS
