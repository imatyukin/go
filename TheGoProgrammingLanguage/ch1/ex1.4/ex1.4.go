// Упражнение 1.4.
// Измените программу dup2 так, чтобы она выводила имена всех файлов, в которых найдены повторяющиеся строки.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileIn := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4: %v\n", err)
				continue
			}
			countLines(f, counts, fileIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, fileIn[line], line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !in(f.Name(), fileIn[line]) {
			fileIn[line] = append(fileIn[line], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func in(value string, strings []string) bool {
	for _, item := range strings {
		if value == item {
			return true
		}
	}
	return false
}