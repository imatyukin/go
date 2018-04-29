// Упражнение 3.10.
// Напишите нерекурсивную версию функции comma, использующую bytes. Buffer вместо конкатенации строк.
package main

import (
	"fmt"
	"os"
	"bytes"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	b := &bytes.Buffer{}
	pre := len(s) % 3
	// Write the first group of up to 3 digits.
	if pre == 0 {
		pre = 3
	}
	b.WriteString(s[:pre])
	// Deal with the rest.
	for i := pre; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}

//!-