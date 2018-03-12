// Упражнение 2.2.
// Напишите программу общего назначения для преобразования единиц, аналогичную cf, которая считывает числа из
// аргументов командной строки (или из стандартного ввода, если аргументы командной строки отсутствуют) и преобразует
// каждое число в другие единицы, как температуру — в градусы Цельсия и Фаренгейта, длину — в футы и метры, вес — в
// фунты и килограммы и т.д.

//!+

// Package tempconv performs Celsius and Fahrenheit, Feet and Meter, Pound and Kilogram conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Feet float64
type Meter float64

type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g meter", m) }
func (p Pound) String() string      { return fmt.Sprintf("%g pound", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%g kilogram", k) }

//!-
