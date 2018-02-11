package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var index = 0
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(index, os.Args[i])
		index++
	}

	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}

	index = 0
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(index, strings.Join(os.Args[i:i+1], "\n"))
		index++
	}
}