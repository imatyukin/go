// Упражнение 2.2.
// Напишите программу общего назначения для преобразования единиц, аналогичную cf, которая считывает числа из
// аргументов командной строки (или из стандартного ввода, если аргументы командной строки отсутствуют) и преобразует
// каждое число в другие единицы, как температуру — в градусы Цельсия и Фаренгейта, длину — в футы и метры, вес — в
// фунты и килограммы и т.д.

//!+

// ex2.2 converts its numeric argument to Celsius, Fahrenheit, Feet, Meter, Pound, Kilogram.
package main

import (
	"fmt"
	"os"
	"strconv"

	"../tempconv1"
)

func process(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex2.2: %v\n", err)
		os.Exit(1)
	}
	{
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	{
		f := tempconv.Feet(t)
		m := tempconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToM(f), m, tempconv.MToF(m))
	}
	{
		p := tempconv.Pound(t)
		k := tempconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n", p, tempconv.PToK(p), k, tempconv.KToP(k))
	}

}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			process(arg)
		}
		return
	}

	fmt.Println("Input number. Ctrl-C to quit.")

	for true {
		var arg string
		_, err := fmt.Scanf("%s", &arg)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		process(arg)
	}
}

//!-
